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

  test "driver location update" do
    conn = conn(:patch, "/drivers/42")

    conn = Router.call(conn, @opts)

    assert get_resp_header(conn, "content-type") == ["application/json; charset=utf-8"]

    assert conn.state == :sent
    assert conn.status == 200
    assert conn.resp_body == "\{\"driver\":\"42\"\}"
  end

  test "driver zombie status" do
    conn = conn(:get, "/drivers/42")

    conn = Router.call(conn, @opts)

    assert get_resp_header(conn, "content-type") == ["application/json; charset=utf-8"]

    assert conn.state == :sent
    assert conn.status == 200
    assert conn.resp_body == "\{\"id\":\"42\"\,\"zombie\":\"true\"}"
  end
end
