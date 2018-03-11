defmodule Gateway.Location do
  @moduledoc """
  Specification for the Location Service NSQ producer.
  """

  @typep driver_id :: integer
  @typep location :: %{latitude: number, longitude: number}

  @doc """
  Updates the location of a given driver identified by `driver_id`.
  """
  @callback publish_location(id :: driver_id, l :: location) :: {:ok, term} | {:error, term}
end
