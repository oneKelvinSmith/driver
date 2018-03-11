defmodule Gateway.Router do
  use Plug.Router

  plug(
    Plug.Parsers,
    parsers: [:urlencoded, :json],
    pass: ["text/*"],
    json_decoder: Poison
  )

  plug(:match)
  plug(:dispatch)

  get "/" do
    conn
    |> set_json_content_type
    |> send_resp(200, "Gateway")
  end

  patch "/drivers/:id" do
    conn
    |> set_json_content_type
    |> update_driver_location(id)
  end

  get "/drivers/:id" do
    conn
    |> set_json_content_type
    |> fetch_zombie_status(id)
  end

  match _ do
    conn
    |> set_json_content_type
    |> send_resp(404, error_response("Not Found"))
  end

  defp update_driver_location(%Plug.Conn{private: private} = conn, driver_id) do
    producer = private[:location_producer] || Gateway.Location.Producer

    case producer.publish_location(driver_id, conn.body_params) do
      {:ok, _} ->
        send_resp(conn, 204, "")

      {:error, _} ->
        message = "Unable to update location of driver: #{driver_id}"
        send_resp(conn, 503, error_response(message))
    end
  end

  defp fetch_zombie_status(%Plug.Conn{private: private} = conn, driver_id) do
    api = private[:zombie_api] || Gateway.Zombie.Api

    case api.status(driver_id) do
      {:ok, zombie_status} ->
        send_resp(conn, 200, Poison.encode!(zombie_status))

      {:error, _} ->
        message = "Unable to retrieve zombie status for driver: #{driver_id}"
        send_resp(conn, 503, error_response(message))
    end
  end

  def set_json_content_type(conn) do
    put_resp_content_type(conn, "application/json", "utf-8")
  end

  defp error_response(message) do
    Poison.encode!(%{error: message})
  end
end
