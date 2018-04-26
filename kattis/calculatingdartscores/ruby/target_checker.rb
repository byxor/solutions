require './converter'

def check(target, a, b, c)
  return (convert(a) + convert(b) + convert(c)) == target
end
