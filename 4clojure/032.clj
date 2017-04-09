(fn duplicate [collection]
  (if (not (= collection []))
    (concat
      (list (first collection) (first collection))
      (duplicate (rest collection)))))
