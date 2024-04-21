package withdraw
import "errors"

func WithDraw(sb, monto int) (int, error) {

	if sb < monto {
		return 0, errors.New("error")
	}

	dinero := sb - monto
	return dinero, nil

}
