defmodule Gateway.Location.Producer do
  @moduledoc """
  Implimentation for the Location Service NSQ Producer.
  """

  use GenServer

  @behaviour Gateway.Location

  @topic "driver"
  @config %NSQ.Config{
    nsqds: ["127.0.0.1:4150"],
    user_agent: "Gateway Producer"
  }

  def init(args) do
    {:ok, args}
  end

  def start_link(opts \\ []) do
    NSQ.Producer.Supervisor.start_link(@topic, @config, opts)
  end

  def update_location(driver_id) do
    NSQ.Producer.pub(__MODULE__, "driver location update #{driver_id}")

    {:ok, {:updated, driver_id}}
  end
end
