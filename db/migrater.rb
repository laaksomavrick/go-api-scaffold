
# eventually allow args for environment and opts
def execute_script(file_name)
  (`psql "dbname='goals_development' user='postgres' password='postgres'" -h 127.0.0.1 -p 5432 -a -f db/migrations/#{file_name}`)
end

abs_path = File.expand_path(File.dirname(__FILE__))
file_names = Dir.children("#{abs_path}/migrations")

file_names.each do |file_name|
  p "Executing #{file_name}"
  execute_script(file_name)
end

# here for if/when keeping track of the last ran migration matters
#open("last_migration", "w") { |f| f << some_file_name }
#open("last_migration", "r") { |f| p f.read() }
