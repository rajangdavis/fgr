require "csv"

# This is a a reference to the file so that it can be read
@data = File.new("../clean_train.csv", "r")

# Initializes a CSV representation for Ruby
@csv = CSV.new(@data, :headers => true, :converters => :all)

# Parses the CSV to the @accounts_to_create array as JSON object
@csv.to_a.map {|row| ret = row.to_hash;  ret.delete(ret.keys.first); ret.slice("Overall Qual","Year Built","Overall Cond", "SalePrice","Year Remod/Add", "Yr Sold") }.to_json

# Send to DOM

## Need to rewrite the above into JS
## create a JS work flow for D3