package runrace

/*
 * Crea una función que evalúe si un/a atleta ha superado correctamente una
 * carrera de obstáculos.
 * - La función recibirá dos parámetros:
 *      - Un array que sólo puede contener String con las palabras
 *        "run" o "jump"
 *      - Un String que represente la pista y sólo puede contener "_" (suelo)
 *        o "|" (valla)
 * - La función imprimirá cómo ha finalizado la carrera:
 *      - Si el/a atleta hace "run" en "_" (suelo) y "jump" en "|" (valla)
 *        será correcto y no variará el símbolo de esa parte de la pista.
 *      - Si hace "jump" en "_" (suelo), se variará la pista por "x".
 *      - Si hace "run" en "|" (valla), se variará la pista por "/".
 * - La función retornará un Boolean que indique si ha superado la carrera.
 * Para ello tiene que realizar la opción correcta en cada tramo de la pista.
 */
import "slices"

const (
	_run    = "run"
	_jump   = "jump"
	_ground = "_"
	_wall   = "|"
)

func Execute(jumpRun []string, pista []string) bool {
	for i, action := range jumpRun {
		step := pista[i]
		isValidAction := validateSteps(action, step)
		if !isValidAction {
			pista[i] = updatePista(action)
		}
	}

	return !slices.Contains(pista, "x") && !slices.Contains(pista, "/")
}

func validateSteps(action, step string) bool {
	if action == _run && step == _ground {
		return true
	}

	if action == _jump && step == _wall {
		return true
	}

	return false
}

func updatePista(action string) string {
	if action == _run {
		return "/"
	}
	return "x"
}
