(fn get_length [collection]
  (if (= collection [])
    0
    (+ (get_length (rest collection)) 1)))
