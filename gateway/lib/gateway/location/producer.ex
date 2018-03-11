defmodule Gateway.Location.Producer do
  @moduledoc """
  Implimentation for the Location Service NSQ Producer.
  """

  @behaviour Gateway.Zombie

  @nsqlookupd "localhost:4161"
  @user_agent [{"User-agent", "Gateway"}]

  def update_location(driver_id) do
    {:ok, {:updated, driver_id}}
  end
end
