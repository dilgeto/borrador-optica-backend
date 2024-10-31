package data

type Empleado struct {
	Persona
	Id_empleado uint32
	Contrasena  string
	Rol         string
}
