(fn palindrome? [collection]
  (= (seq collection) (reverse collection)))
