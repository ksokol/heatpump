package metric

import (
	"fmt"
	"heatpump/internal/heatpump"
)

func Parse(data *heatpump.Data) *Metric {
	m := New()

	parseParameters(data.Parameters, m)
	parseValues(data.Values, m)
	parseState(data, m)

	return m
}

func parseParameters(values *[]int32, m *Metric) {
	for i, value := range *values {
		switch i {
		case 1:
			m.addTemperature("temperature_tolerance", parseDoubleValue(value))
		case 3:
			m.addOperationMode("heating_circuit", value)
		case 4:
			m.addOperationMode("hot_water", value)
		case 105:
			m.addHotWater("temperature_tolerance", parseDoubleValue(value)) //degrees
		case 108:
			m.addCoolingOperationMode("cooling", value)
		case 110:
			m.addCooling("outdoor_temperature_clearance", parseDoubleValue(value)) //degrees
		case 119:
			m.addOperationMode("pool", value)
		case 124:
			m.addHeating("solar_t_diff", parseDoubleValue(value)) //degress
		case 132:
			m.addCooling("set_point_mk1", parseDoubleValue(value)) //degrees
		case 133:
			m.addCooling("set_point_mk2", parseDoubleValue(value)) //degrees
		case 134:
			m.addCooling("working_temperature_difference_1", parseDoubleValue(value)) //kelvin
		case 135:
			m.addCooling("working_temperature_difference_2", parseDoubleValue(value)) //kelvin
		case 696:
			m.addOperationMode("heating_mixed_circuit_2", value)
		case 779:
			m.addOperationMode("heating_mixed_circuit_3", value)
		case 850:
			m.addCooling("outdoor_temperature_overrun", parseDoubleValue(value)) //hours
		case 851:
			m.addCooling("outdoor_temperature_underrun", parseDoubleValue(value)) //hours
		case 881:
			m.addOperationMode("solar", value)
		case 894:
			m.addVentilationOperationMode("ventilation", value)
		case 966:
			m.addCooling("set_point_mk3", parseDoubleValue(value)) //degrees
		case 967:
			m.addCooling("working_temperature_difference_3", parseDoubleValue(value)) //kelvin
		}
	}
}

func parseValues(values *[]int32, m *Metric) {
	for i, value := range *values {
		switch i {
		case 10:
			m.addTemperature("outgoing", parseDoubleValue(value)) //degress
		case 11:
			m.addTemperature("incoming_effective", parseDoubleValue(value)) //degress
		case 12:
			m.addTemperature("incoming_nominal", parseDoubleValue(value)) //degress
		case 13:
			m.addTemperature("incoming_external", parseDoubleValue(value)) //degress
		case 14:
			m.addTemperature("hot_gas", parseDoubleValue(value)) //degress
		case 15:
			m.addTemperature("outdoor", parseDoubleValue(value)) //degress
		case 16:
			m.addTemperature("outdoor_avg", parseDoubleValue(value)) //degress
		case 17:
			m.addTemperature("hot_water_effective", parseDoubleValue(value)) //degress
		case 18:
			m.addTemperature("hot_water_nominal", parseDoubleValue(value)) //degress
		case 19:
			m.addTemperature("probe_in", parseDoubleValue(value)) //degress
		case 20:
			m.addTemperature("probe_out", parseDoubleValue(value)) //degress
		case 21:
			m.addTemperature("mk1", parseDoubleValue(value)) //degress
		case 24:
			m.addTemperature("mk2", parseDoubleValue(value)) //degress
		case 26:
			m.addSolar("solar_collector", parseDoubleValue(value)) //degress
		case 27:
			m.addSolar("solar_tank", parseDoubleValue(value)) //degress
		case 28:
			m.addTemperature("external_source", parseDoubleValue(value)) //degress
		case 37:
			m.addOutput("AV", value)
		case 38:
			m.addOutput("BUP", value)
		case 39:
			m.addOutput("HUP", value)
		case 40:
			m.addOutput("MA1", value)
		case 41:
			m.addOutput("MZ1", value)
		case 42:
			m.addOutput("VEN", value)
		case 43:
			m.addOutput("VBO", value)
		case 44:
			m.addOutput("VD1", value)
		case 45:
			m.addOutput("VD2", value)
		case 46:
			m.addOutput("ZIP", value)
		case 47:
			m.addOutput("ZUP", value)
		case 48:
			m.addOutput("ZW1", value)
		case 49:
			m.addOutput("ZW2SST", value)
		case 50:
			m.addOutput("ZW3SST", value)
		case 51:
			m.addOutput("FP2", value)
		case 52:
			m.addOutput("SLP", value)
		case 53:
			m.addOutput("SUP", value)
		case 54:
			m.addOutput("MZ2", value)
		case 55:
			m.addOutput("MA2", value)
		case 56:
			m.addOperatingHours("cmp1", parseHours(value)) //hours
		case 57:
			m.addOperatingHours("pulse_cmp1", value) //counter

			if !((*values)[56] == 0 || (*values)[57] == 0) {
				m.addOperatingHours(
					"average_runtime_cmp1",
					formatDouble(float32((*values)[56])/float32((*values)[57]*3600)),
				) //hours
			}
		case 58:
			m.addOperatingHours("cmp2", parseHours(value)) //hours
		case 59:
			m.addOperatingHours("pulse_cmp2", parseHours(value)) //counter

			if !((*values)[58] == 0 || (*values)[59] == 0) {
				m.addOperatingHours(
					"average_runtime_cmp2",
					formatDouble(float32((*values)[58])/float32((*values)[59]*3600)),
				) //hours
			}
		case 60:
			m.addOperatingHours("shg1" /* zwe1 */, parseHours(value))
			//float32(value) / 3600) //hours
		case 61:
			m.addOperatingHours("shg2" /* zwe2 */, parseHours(value)) // //hours
		case 62:
			m.addOperatingHours("shg3" /* zwe3 */, parseHours(value)) // //hours
		case 63:
			m.addOperatingHours("heatpump", parseHours(value)) // //hours
		case 64:
			m.addOperatingHours("heating", parseHours(value)) // //hours
		case 65:
			m.addOperatingHours("bw", parseHours(value)) // //hours
		case 66:
			m.addOperatingHours("cooling", parseHours(value)) // //hours
		case 100:
			if value != 0 {
				m.addFault("code", value, int64((*values)[i-5]))
			}
		case 101:
			if value != 0 {
				m.addFault("code", value, int64((*values)[i-5]))
			}
		case 102:
			if value != 0 {
				m.addFault("code", value, int64((*values)[i-5]))
			}
		case 103:
			if value != 0 {
				m.addFault("code", value, int64((*values)[i-5]))
			}
		case 104:
			if value != 0 {
				m.addFault("code", value, int64((*values)[i-5]))
			}
		case 106:
			if value != 0 {
				m.addShutdowns("code", value, int64((*values)[i+5]))
			}
		case 107:
			if value != 0 {
				m.addShutdowns("code", value, int64((*values)[i+5]))
			}
		case 108:
			if value != 0 {
				m.addShutdowns("code", value, int64((*values)[i+5]))
			}
		case 109:
			if value != 0 {
				m.addShutdowns("code", value, int64((*values)[i+5]))
			}
		case 110:
			if value != 0 {
				m.addShutdowns("code", value, int64((*values)[i+5]))
			}
		case 138:
			m.addOutput("MZ3", value)
		case 139:
			m.addOutput("MA3", value)
		case 140:
			m.addOutput("FP3", value)
		case 145:
			m.addOperatingHours("sw", parseHours(value)) //hours
		case 151:
			m.addThermal("heating", parseThermal(value))
		case 152:
			m.addThermal("hot_water", parseThermal(value))
		case 153:
			m.addThermal("pool", parseThermal(value))
		case 154:
			heating := parseThermal((*values)[151])
			hotWater := parseThermal((*values)[152])
			pool := parseThermal((*values)[153])
			m.addThermal("total", heating+hotWater+pool)
		case 155:
			m.addThermal("massflow", value)
		case 161:
			m.addOperatingHours("solar", parseHours(value)) //hours
		case 166:
			m.addOutput("VSK", value)
		case 167:
			m.addOutput("FRH", value)
		case 213:
			m.addOutput("AV2", value)
		case 214:
			m.addOutput("VBO2", value)
		case 215:
			m.addOutput("VD12", value)
		case 216:
			m.addOutput("VDH2", value)
		}
	}
}

func parseState(data *heatpump.Data, m *Metric) {
	parameters := data.Parameters
	values := data.Values

	var ventilationState int32
	var coolingState int32
	var hotWaterState int32
	var heatingState int32
	var poolState int32

	if (*values)[80] == 0 {
		heatingState = 2
	} else if (*parameters)[3] == 4 {
		heatingState = 0
	} else {
		var v1 int32

		if (*parameters)[35] != 0 {
			v1 = (*values)[11]
		} else {
			v1 = (*values)[13]
		}

		v2 := (*values)[12] - (*parameters)[88]

		if v1 < v2 {
			heatingState = 1
		} else {
			heatingState = 3
		}
	}

	m.addState("heating", heatingState)

	param := (*values)[17]

	if (*values)[80] == 1 {
		hotWaterState = 2
	} else if (*parameters)[4] == 4 {
		hotWaterState = 0
	} else if (*parameters)[82] == 1 {
		if param > 0 {
			hotWaterState = 1
		} else {
			hotWaterState = 3
		}
	} else if param < (*values)[18]-(*parameters)[74] {
		hotWaterState = 1
	} else {
		hotWaterState = 3
	}

	m.addState("water", hotWaterState)

	if (*values)[80] == 2 {
		poolState = 2
	} else if (*parameters)[119] == 4 {
		poolState = 0
	} else if (*values)[36] > 0 {
		poolState = 1
	} else {
		poolState = 3
	}

	m.addState("pool", poolState)

	if (*values)[80] == 7 {
		coolingState = 2
	} else if (*parameters)[108] == 0 {
		coolingState = 0
	} else {
		coolingState = 3
	}
	m.addState("cooling", coolingState)

	if (*parameters)[894] != 3 {
		ventilationState = 3
	} else {
		ventilationState = 0
	}

	m.addState("ventilation", ventilationState)
	m.addState("solar", 0)
}

func parseDoubleValue(v int32) string {
	return formatDouble(float32(v) / 10)
}

func parseHours(v int32) float32 {
	return float32(v) / 3600
}

func formatDouble(v float32) string {
	return fmt.Sprintf("%.1f", v)
}

// workaround for thermal energies
// the thermal energies can be unreasonably high in some cases,
// probably due to a sign bug in the firmware
// trying to correct this issue here
func parseThermal(v int32) int32 {
	if v >= 214748364 {
		v = v - 214748364
	}
	return v / 10
}
