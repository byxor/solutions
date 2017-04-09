(fn sum [numbers]
  (if (= numbers [])
    0
    (+
      (first numbers)
      (sum (rest numbers)))))
