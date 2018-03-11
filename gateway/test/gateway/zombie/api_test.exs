defmodule Gateway.Zombie.ApiTest do
  use ExUnit.Case
  doctest Gateway

  alias Gateway.Zombie.Api

  @tag :httpoison
  describe "status/1" do
    setup do
      HTTPoison.start
      :ok
    end

    test "returns the cleaned up response from HTTPoison get" do
      assert Api.status(42) == {:ok, %{"id" => 42, "zombie" => true}}
    end
  end

  describe "handle_response/1" do
    test "returns the body of a successful response as on :ok tuple" do
      response = {:ok, %{status_code: 200, body: "[{\"key\":\"value\"}]"}}
      assert Api.handle_response(response) == {:ok, [%{"key" => "value"}]}
    end

    test "returns the body of an unsuccessful response as on :error tuple" do
      response = {:anything, %{status_code: 999, body: "[{\"key\":\"value\"}]"}}
      assert Api.handle_response(response) == {:error, [%{"key" => "value"}]}
    end
  end
end
