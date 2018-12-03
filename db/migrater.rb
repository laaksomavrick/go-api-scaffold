
# eventually allow args for environment and opts
# todo clean this up - migrate to go for test spinup/teardown?
def execute_script(file_name, i, dbname = "goals_development")
  if i == 0
    (`psql "user='postgres' password='postgres'" -h 127.0.0.1 -p 5432 -a -f db/migrations/#{file_name}`)
  else
    (`psql "dbname='#{dbname}' user='postgres' password='postgres'" -h 127.0.0.1 -p 5432 -a -f db/migrations/#{file_name}`)
  end
end

abs_path = File.expand_path(File.dirname(__FILE__))
file_names = Dir.children("#{abs_path}/migrations")

file_names.sort.each_with_index do |file_name, i|
  p "Executing #{file_name}"
  execute_script(file_name, i)
end

# here for if/when keeping track of the last ran migration matters
#open("last_migration", "w") { |f| f << some_file_name }
#open("last_migration", "r") { |f| p f.read() }
