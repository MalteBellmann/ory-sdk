=begin
#Ory APIs

#Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

The version of the OpenAPI document: v1.6.1
Contact: support@ory.sh
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 6.0.1

=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for OryClient::MigrationOptions
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe OryClient::MigrationOptions do
  let(:instance) { OryClient::MigrationOptions.new }

  describe 'test an instance of MigrationOptions' do
    it 'should create an instance of MigrationOptions' do
      expect(instance).to be_instance_of(OryClient::MigrationOptions)
    end
  end
  describe 'test attribute "environment"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
      # validator = Petstore::EnumTest::EnumAttributeValidator.new('String', ["prod", "dev"])
      # validator.allowable_values.each do |value|
      #   expect { instance.environment = value }.not_to raise_error
      # end
    end
  end

  describe 'test attribute "project_subscription"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
      # validator = Petstore::EnumTest::EnumAttributeValidator.new('String', ["migrate", "ignore"])
      # validator.allowable_values.each do |value|
      #   expect { instance.project_subscription = value }.not_to raise_error
      # end
    end
  end

end
