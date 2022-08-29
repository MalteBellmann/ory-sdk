=begin
#Ory APIs

#Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

The version of the OpenAPI document: v0.2.0-alpha.22
Contact: support@ory.sh
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 6.0.1

=end

require 'date'
require 'time'

module OryClient
  class NullPlan
    UNKNOWN = "unknown".freeze
    FREE = "free".freeze
    START_UP_MONTHLY = "start_up_monthly".freeze
    START_UP_YEARLY = "start_up_yearly".freeze
    BUSINESS_MONTHLY = "business_monthly".freeze
    BUSINESS_YEARLY = "business_yearly".freeze
    CUSTOM = "custom".freeze

    # Builds the enum from string
    # @param [String] The enum value in the form of the string
    # @return [String] The enum value
    def self.build_from_hash(value)
      new.build_from_hash(value)
    end

    # Builds the enum from string
    # @param [String] The enum value in the form of the string
    # @return [String] The enum value
    def build_from_hash(value)
      constantValues = NullPlan.constants.select { |c| NullPlan::const_get(c) == value }
      raise "Invalid ENUM value #{value} for class #NullPlan" if constantValues.empty?
      value
    end
  end
end
