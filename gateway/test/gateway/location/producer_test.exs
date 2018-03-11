defmodule Gateway.Location.ProducerTest do
  use ExUnit.Case, async: true
  doctest Gateway

  alias Gateway.Location.Producer

  setup do
    producer = start_supervised!(Producer)
    %{producer: producer}
  end

  @tag :nsq
  test "exists" do
    driver_id = 42

    assert Producer.update_location(driver_id) == {:ok, "OK"}
  end
end
