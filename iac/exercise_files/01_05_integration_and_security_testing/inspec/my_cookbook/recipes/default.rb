#
# Cookbook:: my_cookbook
# Recipe:: default
#
# Copyright:: 2017, The Authors, All Rights Reserved.
#

#nodes = search(:node, "hostname:[* TO *] AND chef_environment:#{node.chef_environment}")

package 'apache2' do
  action :install
end
