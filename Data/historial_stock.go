package data

type Historial_stock struct {
	Id_historial      uint32
	Fecha_cambio      uint32
	Tipo_cambio       string
	Cantidad_cambiada int32
	Stock_anterior    int32
	Stock_nuevo       int32
	Detalle           string
}
