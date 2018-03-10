defmodule Gateway.Router do
  use Plug.Router

  plug(:match)
  plug(:dispatch)

  get "/" do
    conn
    |> put_resp_content_type("application/json", "utf-8")
    |> send_resp(200, "Gateway")
  end

  match _ do
    conn
    |> put_resp_content_type("application/json", "utf-8")
    |> send_resp(404, "\{\"error\":\"Not Found\"\}")
  end
end
