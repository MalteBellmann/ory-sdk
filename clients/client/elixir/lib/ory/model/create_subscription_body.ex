# NOTE: This file is auto generated by OpenAPI Generator 7.2.0 (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule Ory.Model.CreateSubscriptionBody do
  @moduledoc """
  Create Subscription Request Body
  """

  @derive Jason.Encoder
  defstruct [
    :currency,
    :interval,
    :plan,
    :provision_first_project,
    :return_to,
    :workspace
  ]

  @type t :: %__MODULE__{
    :currency => String.t | nil,
    :interval => String.t,
    :plan => String.t,
    :provision_first_project => String.t | nil,
    :return_to => String.t | nil,
    :workspace => String.t | nil
  }

  def decode(value) do
    value
  end
end

