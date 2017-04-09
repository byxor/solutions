(fn get_nth [collection index]
  (if (= index 0)
    (first collection)
    (recur (rest collection) (- index 1))))
