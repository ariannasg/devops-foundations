# -*- mode: ruby -*-
# # vi: set ft=ruby :

Vagrant.configure("2") do |config|

  config.vm.define "web" do |web|
    web.vm.box = "bento/ubuntu-14.04"
    web.vm.network "forwarded_port", guest: 80, host: 8080
    web.vm.synced_folder "./html", "/var/www/html/class"
    web.vm.provision :shell, path: "bootstrap.sh"
  end

  config.vm.define "db" do |db|
    db.vm.box = "bento/ubuntu-12.04"
    db.vm.network "forwarded_port", guest: 3306, host: 3306
    db.vm.hostname = "dbserver"
    db.vm.provision :shell, path: "db-bootstrap.sh"
  end

end
