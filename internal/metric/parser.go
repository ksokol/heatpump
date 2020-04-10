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
			m.addTemperature("heating_temperature_tolerance", parseDoubleValue(value))
		case 4:
			m.addOperationMode("hot_water", value)
		case 74:
			m.addThermal("hysteresis_hot_water", parseDoubleValue(value))
		case 124:
			m.addTemperature("solar_t_diff", parseDoubleValue(value))
		case 779:
			m.addOperationMode("heating_mixed_circuit_3", value)
		case 881:
			m.addOperationMode("solar", value)
		}
	}
}

func parseValues(values *[]int32, m *Metric) {
	for i, value := range *values {
		switch i {
		case 10:
			m.addTemperature("outgoing", parseDoubleValue(value)) //Vorlauf
		case 11:
			m.addTemperature("incoming_effective", parseDoubleValue(value)) //Rücklauf
		case 12:
			m.addTemperature("incoming_nominal", parseDoubleValue(value)) //Rücklauf-Soll
		case 13:
			m.addTemperature("incoming_external", parseDoubleValue(value)) // Externe Energ.Quelle
		case 14:
			m.addTemperature("hot_gas", parseDoubleValue(value)) //Heissgas
		case 15:
			m.addTemperature("outdoor", parseDoubleValue(value)) //Außentemperatur
		case 16:
			m.addTemperature("outdoor_avg", parseDoubleValue(value)) //Mittelwertemperatur
		case 17:
			m.addTemperature("hot_water_effective", parseDoubleValue(value)) //Warmwasser-Ist
		case 18:
			m.addTemperature("hot_water_max_temperature", parseDoubleValue(value)) //Warmwasser-Soll
		case 19:
			m.addTemperature("probe_in", parseDoubleValue(value))
		case 26:
			m.addTemperature("solar_collector", parseDoubleValue(value))
		case 27:
			m.addTemperature("solar_tank", parseDoubleValue(value))
		case 28:
			m.addTemperature("external_source", parseDoubleValue(value))
		case 37:
			m.addOutput("AV", value) //Abtauventil
		case 38:
			m.addOutput("BUP", value) //Brauchwasserpumpe/Umstellventil/Trinkwasserumwälzpumpe
		case 39:
			m.addOutput("HUP", value) //Heizungsumwälzpumpe
		case 40:
			m.addOutput("MA1", value) //Mischkreis 1 auf
		case 41:
			m.addOutput("MZ1", value) //Mischkreis 1 zu
		case 42:
			m.addOutput("VEN", value) //Ventilation/Lüftung
		case 43:
			m.addOutput("VBO", value) //Solepumpe/Ventilator
		case 44:
			m.addOutput("VD1", value) //Verdichter 1
		case 47:
			m.addOutput("ZUP", value) //Zusatzumwälzpumpe
		case 48:
			m.addOutput("ZW1", value) //Steuersignal Zusatzheizung v. Heizung
		case 51:
			m.addOutput("FP2", value) //Pumpe Mischkreis 2 TODO nachfragen
		case 52:
			m.addOutput("SLP", value) //Solarladepumpe
		case 54:
			m.addOutput("MZ2", value) //Mischkreis 2 zu TODO nachfragen
		case 55:
			m.addOutput("MA2", value) //Mischkreis 2 auf TODO nachfragen
		case 56:
			m.addOperatingHour("cmp1", value)
		case 57:
			m.addCount("pulse_cmp1", value)

			if !((*values)[56] == 0 || (*values)[57] == 0) {
				m.addOperatingHour(
					"cmp1_runtime_avg",
					formatDouble(float32((*values)[56]/60)/float32((*values)[57])),
				)
			}
		case 60:
			m.addOperatingHour("shg1", value) //zweiter Wärmerzeuger 1
		case 64:
			m.addOperatingHour("heating", value)
		case 65:
			m.addOperatingHour("hot_water", value)
		case 80:
			m.addOperationState("heatpump", value)
		case 88:
			m.addThermal("hysteresis_heating", parseDoubleValue(value))
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
				m.addShutdown("code", value, int64((*values)[i+5]))
			}
		case 107:
			if value != 0 {
				m.addShutdown("code", value, int64((*values)[i+5]))
			}
		case 108:
			if value != 0 {
				m.addShutdown("code", value, int64((*values)[i+5]))
			}
		case 109:
			if value != 0 {
				m.addShutdown("code", value, int64((*values)[i+5]))
			}
		case 110:
			if value != 0 {
				m.addShutdown("code", value, int64((*values)[i+5]))
			}
		case 138:
			m.addOutput("MZ3", value) //Mischkreis 3 zu TODO nachfragen
		case 139:
			m.addOutput("MA3", value) //Mischkreis 3 auf TODO nachfragen
		case 147:
			m.addVolt("AIn", parseVolt(value)) //Analoges Eingangssignal
		case 151:
			m.addEnergy("heating", parseEnergy(value))
		case 152:
			m.addEnergy("hot_water", parseEnergy(value))
		case 157:
			m.addVolt("AO2", parseVolt(value)) //Analoges Ausgangssignal 2
		case 161:
			m.addOperatingHour("solar", value)
		case 166:
			m.addOutput("VSK", value) //TODO nachfragen
		case 167:
			m.addOutput("FRH", value) //TODO nachfragen
		case 176:
			m.addTemperature("AVD", parseDoubleValue(value)) //Ansaug VD
		case 177:
			m.addTemperature("VDH", parseDoubleValue(value)) //VD-Heizung
		case 178:
			m.addThermal("UE", parseDoubleValue(value)) //Überhitzung
		case 180:
			m.addPressure("HD", parseDoubleValue(value))
		case 181:
			m.addPressure("ND", parseDoubleValue(value))
		case 213:
			m.addOutput("AV2", value) //Abtauventil 2 TODO nachfragen
		case 214:
			m.addOutput("VBO2", value) //Solepumpe/Ventilator TODO nachfragen
		case 215:
			m.addOutput("VD12", value) //Verdichter 1/2  TODO nachfragen
		case 216:
			m.addOutput("VDH2", value) //Verdichterheizung 2 TODO nachfragen
		}
	}
}

func parseState(data *heatpump.Data, m *Metric) {
	parameters := data.Parameters
	values := data.Values

	var hotWaterState int32
	var heatingState int32

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

	if (*values)[80] == 1 {
		hotWaterState = 2
	} else if (*parameters)[4] == 4 {
		hotWaterState = 0
	} else if (*parameters)[82] == 1 {
		if (*values)[17] > 0 {
			hotWaterState = 1
		} else {
			hotWaterState = 3
		}
	} else if (*values)[17] < (*values)[18]-(*parameters)[74] {
		hotWaterState = 1
	} else {
		hotWaterState = 3
	}

	m.addState("water", hotWaterState)
}

func parseDoubleValue(v int32) string {
	return formatDouble(float32(v) / 10)
}

func parseVolt(v int32) string {
	return fmt.Sprintf("%.2f", float32(v)/100)
}

func formatDouble(v float32) string {
	return fmt.Sprintf("%.1f", v)
}

// workaround for thermal energies
// the thermal energies can be unreasonably high in some cases,
// probably due to a sign bug in the firmware
// trying to correct this issue here
func parseEnergy(v int32) int32 {
	if v >= 214748364 {
		v = v - 214748364
	}
	return v / 10
}
