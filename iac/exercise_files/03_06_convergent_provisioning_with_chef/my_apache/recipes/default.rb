#
# Cookbook:: my_apache
# Recipe:: default
#
# Copyright:: 2017, The Authors, All Rights Reserved.
#

# package 'apache2'

package 'Install Apache' do
  case node[:platform]
  when 'redhat', 'centos'
    package_name 'httpd'
  when 'ubuntu', 'debian'
    package_name 'apache2'
  end
end
