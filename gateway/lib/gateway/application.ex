defmodule Gateway.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  alias Plug.Adapters.Cowboy

  def start(_type, _args) do
    children = [
      Cowboy.child_spec(:http, Gateway.Router, [], port: 3000)
    ]

    opts = [strategy: :one_for_one, name: Gateway.Supervisor]
    Supervisor.start_link(children, opts)
  end
end
