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
    coordinates = %{}
    response = Producer.publish_location(42, coordinates)

    assert response == {:ok, "OK"}
  end
end
