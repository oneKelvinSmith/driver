defmodule Gateway.RouterTest do
  use ExUnit.Case, async: true
  use Plug.Test

  alias Gateway.Router

  @opts Router.init([])

  test "health check" do
    conn = conn(:get, "/")

    conn = Router.call(conn, @opts)

    assert get_resp_header(conn, "content-type") == ["application/json; charset=utf-8"]

    assert conn.state == :sent
    assert conn.status == 200
    assert conn.resp_body == "Gateway"
  end

  test "404" do
    conn = conn(:get, "/anything_else")

    conn = Router.call(conn, @opts)

    assert get_resp_header(conn, "content-type") == ["application/json; charset=utf-8"]

    assert conn.state == :sent
    assert conn.status == 404
    assert conn.resp_body == "\{\"error\":\"Not Found\"\}"
  end

  describe "driver location update" do
    defmodule LocationProducer do
      @behaviour Gateway.Location

      def publish_location(_driver_id, _location) do
        {:ok, "OK"}
      end
    end

    test "when NSQ is available" do
      update_body =
        Poison.encode!(%{
          latitude: 48.8566,
          longitude: 2.3522
        })

      conn =
        conn(:patch, "/drivers/42", update_body)
        |> put_req_header("content-type", "application/json")
        |> Plug.Conn.put_private(:location_producer, LocationProducer)
        |> Router.call(@opts)

      assert get_resp_header(conn, "content-type") == ["application/json; charset=utf-8"]

      assert conn.state == :sent
      assert conn.status == 204
      assert conn.resp_body == ""
    end

    defmodule UnavailableProducer do
      @behaviour Gateway.Location

      def publish_location(_driver_id, _location) do
        {:error, :unavailable}
      end
    end

    test "when NSQ is NOT available" do
      update_body =
        Poison.encode!(%{
          latitude: 48.8566,
          longitude: 2.3522
        })

      conn =
        conn(:patch, "/drivers/42", update_body)
        |> put_req_header("content-type", "application/json")
        |> Plug.Conn.put_private(:location_producer, UnavailableProducer)
        |> Router.call(@opts)

      assert get_resp_header(conn, "content-type") == ["application/json; charset=utf-8"]

      assert conn.state == :sent
      assert conn.status == 503
      assert conn.resp_body == "\{\"error\":\"Unable to update location of driver: 42\"\}"
    end
  end

  describe "driver zombie status" do
    defmodule ZombieDriverApi do
      @behaviour Gateway.Zombie

      def status(driver_id) do
        {:ok, %{"id" => String.to_integer(driver_id), "zombie" => true}}
      end
    end

    test "when driver is a zombie" do
      conn =
        conn(:get, "/drivers/42")
        |> Plug.Conn.put_private(:zombie_api, ZombieDriverApi)
        |> Router.call(@opts)

      assert get_resp_header(conn, "content-type") == ["application/json; charset=utf-8"]

      assert conn.state == :sent
      assert conn.status == 200
      assert Poison.decode!(conn.resp_body) == %{"id" => 42, "zombie" => true}
    end

    defmodule AliveDriverApi do
      @behaviour Gateway.Zombie

      def status(driver_id) do
        {:ok, %{"id" => String.to_integer(driver_id), "zombie" => false}}
      end
    end

    test "when driver is NOT a zombie" do
      conn =
        conn(:get, "/drivers/42")
        |> Plug.Conn.put_private(:zombie_api, AliveDriverApi)
        |> Router.call(@opts)

      assert get_resp_header(conn, "content-type") == ["application/json; charset=utf-8"]

      assert conn.state == :sent
      assert conn.status == 200
      assert Poison.decode!(conn.resp_body) == %{"id" => 42, "zombie" => false}
    end

    defmodule BrokenZombieApi do
      @behaviour Gateway.Zombie

      def status(_driver_id) do
        {:error, :some_error}
      end
    end

    test "when zombie api has an error" do
      conn =
        conn(:get, "/drivers/42")
        |> Plug.Conn.put_private(:zombie_api, BrokenZombieApi)
        |> Router.call(@opts)

      assert get_resp_header(conn, "content-type") == ["application/json; charset=utf-8"]

      assert conn.state == :sent
      assert conn.status == 503
      assert conn.resp_body == "\{\"error\":\"Unable to retrieve zombie status for driver: 42\"\}"
    end
  end
end
