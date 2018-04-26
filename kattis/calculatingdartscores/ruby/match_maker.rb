require './match_finder'

def make_matches
  matches = []
  (1..180).each { |target| matches << find_matches(target)}
  return matches
end
