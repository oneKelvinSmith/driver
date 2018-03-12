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

  def publish_location(driver_id, location) do
    if service_available() do
      NSQ.Producer.pub(__MODULE__, encode(driver_id, location))
    else
      {:error, :unavailable}
    end
  end

  defp encode(driver_id, location) when is_binary(driver_id) do
    Poison.encode!(%{driver_id: String.to_integer(driver_id), location: location})
  end

  defp encode(driver_id, location) do
    Poison.encode!(%{driver_id: driver_id, location: location})
  end

  defp service_available do
    Enum.any?(Supervisor.which_children(__MODULE__))
  end

  defp topic do
    get_env(:nsqd_topic)
  end

  defp config do
    %NSQ.Config{
      nsqds: [get_env(:nsqd_host)],
      user_agent: @user_agent
    }
  end

  defp get_env(variable) do
    Application.get_env(:gateway, variable)
  end
end
