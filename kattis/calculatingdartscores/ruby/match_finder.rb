require './generator'
require './target_checker'

def trim(array)
  return array.select { |element| element != "single 0" }
end

def find_matches(target)
  if target < 1 or target > 180
    return ["impossible"]
  else
    numbers = generate
    numbers.each { |a|
      numbers.each { |b|
        numbers.each { |c|
          potential = [a, b, c]
          return trim(potential) if check(target, *potential)
        }
      }
    }
    return ["impossible"]
  end
end
