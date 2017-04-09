(fn get_max [& numbers]
  ((fn get_max_ [numbers_]
    (let [head (first numbers_), tail (rest numbers_)]
      (if (= (count numbers_) 1)
        head
        (let [max_of_tail (get_max_ tail)]
          (if (> head max_of_tail)
            head
            max_of_tail)))))
    numbers))
