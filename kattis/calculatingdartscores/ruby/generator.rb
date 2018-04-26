def generate
  numbers = []
  (1..20).each { |i| numbers << "single " + i.to_s }
  (1..20).each { |i| numbers << "double " + i.to_s }
  (1..20).each { |i| numbers << "triple " + i.to_s }
  numbers << "single 0"
  return numbers
end
