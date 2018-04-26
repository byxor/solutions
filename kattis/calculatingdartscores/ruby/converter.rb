def trim_and_multiply(coefficient, word)
  return coefficient * word[7..-1].to_i
end

def convert(word)
  if word.start_with?("single ")
    return trim_and_multiply(1, word)
  elsif word.start_with?("double ")
    return trim_and_multiply(2, word)
  else word.start_with?("triple ")
    return trim_and_multiply(3, word)
  end
end
