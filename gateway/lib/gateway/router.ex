defmodule Gateway.Router do
  use Plug.Router

  plug :match
  plug :dispatch

  get "/" do
    conn
    |> put_resp_content_type("application/json", "utf-8")
    |> send_resp(200, "Gateway")
  end

  patch "/drivers/:id" do
    conn
    |> put_resp_content_type("application/json", "utf-8")
    |> send_resp(200, "\{\"driver\":\"#{id}\"\}")
  end

  get "/drivers/:id" do
    conn
    |> put_resp_content_type("application/json", "utf-8")
    |> fetch_zombie_status(id)
  end

  match _ do
    conn
    |> put_resp_content_type("application/json", "utf-8")
    |> send_resp(404, "\{\"error\":\"Not Found\"\}")
  end

  defp fetch_zombie_status(%Plug.Conn{private: private} = conn, driver_id) do
    zombie_api = private[:zombie_api] || Gateway.Zombie.Api

    case zombie_api.status(driver_id) do
      {:ok, zombie_status} ->
        send_resp(conn, 200, Poison.encode!(zombie_status))
      {:error, _} ->
        send_resp(conn, 400, "Unable to retrieve zombie status for driver: #{driver_id}")
    end
  end
end
