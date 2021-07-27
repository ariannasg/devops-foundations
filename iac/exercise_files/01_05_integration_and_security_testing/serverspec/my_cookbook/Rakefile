require 'rubocop/rake_task'

RuboCop::RakeTask.new(:rubocop) do |t|
  t.options = ['--display-cop-names']
end

desc 'Runs foodcritic against all the cookbooks.'
task :foodcritic do
  sh 'foodcritic .'
end

task default: [:rubocop, :foodcritic]
