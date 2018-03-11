defmodule Gateway.Router do
  use Plug.Router

  plug(:match)
  plug(:dispatch)

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
    |> send_resp(200, "\{\"id\":\"#{id}\"\,\"zombie\":\"true\"}")
  end

  match _ do
    conn
    |> put_resp_content_type("application/json", "utf-8")
    |> send_resp(404, "\{\"error\":\"Not Found\"\}")
  end
end
