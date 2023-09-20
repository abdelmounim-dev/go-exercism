package collatzconjecture
import "errors"

func CollatzConjecture(n int) (steps int, err error) {
  if n <= 0 {
    return 0, errors.New("expected a non-zero positive input")
  }
	for n != 1{
	  if n % 2 == 0 {
      n = n / 2
    } else {
      n = 3 * n + 1
    }
    steps ++
	}
  return steps, nil
}
