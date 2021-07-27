# # encoding: utf-8

# Inspec test for recipe my_cookbook::default

# The Inspec reference, with examples and extensive documentation, can be
# found at http://inspec.io/docs/reference/resources/

describe port(80) do
  it { should be_listening }
end

describe port(3306) do
  it { should_not be_listening }
end
