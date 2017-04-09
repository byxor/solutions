(fn get_odds [collection]
  (if (= collection [])
    (list)
    (if (= (mod (first collection) 2) 1)
       (conj (get_odds (rest collection)) (first collection))
       (get_odds (rest collection)))))
