defmodule Gateway.Location.Producer do
  @moduledoc """
  Implimentation for the Location Service NSQ Producer.
  """

  @behaviour Gateway.Location

  use GenServer

  @user_agent "Gateway Producer/0.42"

  def init(args) do
    {:ok, args}
  end

  def start_link(opts \\ []) do
    NSQ.Producer.Supervisor.start_link(topic(), config(), opts)
  end

  def update_location(driver_id) do
    if Enum.any?(Supervisor.which_children(__MODULE__)) do
      NSQ.Producer.pub(__MODULE__, "driver location update #{driver_id}")
    else
      {:error, :unavailable}
    end
  end

  defp topic do
    get_env(:nsqd_topic)
  end

  defp config do
    port = get_env(:nsqd_port)
    host = get_env(:nsqd_host)

    %NSQ.Config{
      nsqds: [host <> ":" <> port],
      user_agent: @user_agent
    }
  end

  defp get_env(variable) do
    Application.get_env(:gateway, variable)
  end
end
