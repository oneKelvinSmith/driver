defmodule Gateway.Location do
  @moduledoc """
  Specification for the Location Service NSQ producer.
  """

  @typep driver_id :: integer

  @doc """
  Updates the location of a given driver identified by `driver_id`.
  """
  @callback update_location(driver_id :: driver_id) :: {:ok, term} | {:error, term}
end
