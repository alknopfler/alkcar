package main

type Agenda []struct {
	Identrada     string `json:"IdEntrada"`
	NombreEmpresa string `json:"Nombre_Empresa"`
	Letra         string `json:"Letra"`
	Zona          string `json:"Zona"`
	Telefono1     string `json:"Telefono1"`
	Extension1    string `json:"Extension1"`
	Telefono2     string `json:"Telefono2"`
	Extension2    string `json:"Extension2"`
	Telefono3     string `json:"Telefono3"`
	Extension3    string `json:"Extension3"`
	Telefono4     string `json:"Telefono4"`
	Extension4    string `json:"Extension4"`
	Notas         string `json:"Notas"`
	Adr           string `json:"ADR"`
}

type Camiones []struct {
	MatriculacamiN    string `json:"MatriculaCamión"`
	Proveedor         string `json:"Proveedor"`
	Tonelaje          string `json:"Tonelaje"`
	Adr               string `json:"ADR"`
	Observaciones     string `json:"Observaciones"`
	TarjetaTransporte string `json:"Tarjeta_Transporte"`
}

type Clientes []struct {
	Cliente                     string `json:"Cliente"`
	Nif                         string `json:"NIF"`
	Actividad                   string `json:"Actividad"`
	DomicilioCliente            string `json:"Domicilio_Cliente"`
	AmpliacionDomCliente        string `json:"Ampliacion_Dom_Cliente"`
	PoblaciNCliente             string `json:"Población_Cliente"`
	ProvinciaCliente            string `json:"Provincia_Cliente"`
	CodpostCliente              string `json:"CodPost_Cliente"`
	EmpresaFacturar             string `json:"Empresa(Facturar)"`
	DomicilioFacturar           string `json:"Domicilio(Facturar)"`
	AmpliacionDomicilioFacturar string `json:"Ampliacion_Domicilio(Facturar)"`
	PoblaciNFacturar            string `json:"Población(Facturar)"`
	ProvinciaFacturar           string `json:"Provincia(Facturar)"`
	CodpostFacturar             string `json:"CodPost(Facturar)"`
	PersonaDeContacto           string `json:"Persona de contacto"`
	Cargo                       string `json:"Cargo"`
	TelefonoCliente             string `json:"Telefono_Cliente"`
	Extension                   string `json:"Extension"`
	TwoTelefono                 string `json:"2ºTelefono"`
	TwoExtension                string `json:"2ªExtension"`
	ThreeTelefono               string `json:"3ºTelefono"`
	ThreeExtension              string `json:"3ªExtension"`
	FaxCliente                  string `json:"Fax_Cliente"`
	Formadepago                 string `json:"FormadePago"`
	Seguro                      string `json:"Seguro"`
	TarifaCliente               string `json:"Tarifa_Cliente"`
	ObservacionesCliente        string `json:"Observaciones_Cliente"`
}

type ClientePotencial []struct {
	ClientePot                     string `json:"Cliente(Pot)"`
	NifPot                         string `json:"NIF(Pot)"`
	ActividadPot                   string `json:"Actividad(Pot)"`
	DomicilioClientePot            string `json:"Domicilio_Cliente(Pot)"`
	AmpliacionDomClientePot        string `json:"Ampliacion_Dom_Cliente(Pot)"`
	PoblaciNClientePot             string `json:"Población_Cliente(Pot)"`
	ProvinciaClientePot            string `json:"Provincia_Cliente(Pot)"`
	CodpostPot                     string `json:"CodPost(Pot)"`
	EmpresaFacturarPot             string `json:"Empresa(Facturar)(Pot)"`
	DomicilioFacturarPot           string `json:"Domicilio(Facturar)(Pot)"`
	AmpliacionDomicilioFacturarPot string `json:"Ampliacion_Domicilio(Facturar)(Pot)"`
	PoblaciNFacturarPot            string `json:"Población(Facturar)(Pot)"`
	ProvinciaFacturarPot           string `json:"Provincia(Facturar)(Pot)"`
	CodpostFacturarPot             string `json:"CodPost(Facturar)(Pot)"`
	PersonaDeContactoPot           string `json:"Persona de contacto(Pot)"`
	CargoPot                       string `json:"Cargo(Pot)"`
	TelefonoPot                    string `json:"Telefono(Pot)"`
	ExtensionPot                   string `json:"Extension(Pot)"`
	TwoTelefonoPot                 string `json:"2ºTelefono(Pot)"`
	TwoExtensionPot                string `json:"2ªExtension(Pot)"`
	ThreeTelefonoPot               string `json:"3ºTelefono(Pot)"`
	ThreeExtensionPot              string `json:"3ªExtension(Pot)"`
	FaxPot                         string `json:"Fax(Pot)"`
	FormadepagoPot                 string `json:"FormadePago(Pot)"`
	TarifaClientePot               string `json:"Tarifa_Cliente(Pot)"`
	Icluirenclientes               string `json:"IcluirEnClientes"`
	ObservacionesPot               string `json:"Observaciones(Pot)"`
	Estado                         string `json:"Estado"`
}

type ContactoComercialClientes []struct {
	Idcontacto            string `json:"IdContacto"`
	Cliente               string `json:"Cliente"`
	Fechacontacto         string `json:"FechaContacto"`
	AsuntoContacto        string `json:"Asunto_Contacto"`
	ObservacionesContacto string `json:"Observaciones_Contacto"`
}

type ContactoComercialClientesPotenciales []struct {
	IdcontactoPot    string `json:"IdContacto(Pot)"`
	ClientePot       string `json:"Cliente(Pot)"`
	FechacontactoPot string `json:"FechaContacto(Pot)"`
	AsuntoPot        string `json:"Asunto(Pot)"`
	ObservacionesPot string `json:"Observaciones(Pot)"`
}

type Contador []struct {
	Idcontador string `json:"IdContador"`
	Contador   string `json:"Contador"`
}

type ErrorPegados []struct {
	Campo0 string `json:"Campo0"`
}

type Portes []struct {
	Idporte            string `json:"IdPorte"`
	Cliente            string `json:"Cliente"`
	Matricula          string `json:"Matricula"`
	PoblacionOrigen    string `json:"Poblacion_Origen"`
	ProvinciaOrigen    string `json:"Provincia_Origen"`
	PoblacionDestino   string `json:"Poblacion_Destino"`
	ProvinciaDestino   string `json:"Provincia_Destino"`
	Concepto           string `json:"Concepto"`
	SRef               string `json:"S/Ref"`
	FechaCarga         string `json:"Fecha_Carga"`
	NumeroBultos       string `json:"Numero_Bultos"`
	Peso               string `json:"Peso"`
	PrecioPorte        string `json:"Precio_Porte"`
	ConceptoCondEsp    string `json:"Concepto_Cond_Esp"`
	PorteAdr           string `json:"Porte_ADR"`
	ObservacionesPorte string `json:"Observaciones_Porte"`
	Pagar              string `json:"Pagar"`
	FechaPago          string `json:"Fecha_Pago"`
	TalonNumero        string `json:"Talon_Numero"`
	Conforme           string `json:"Conforme"`
	Retenido           string `json:"Retenido"`
	Facturado          string `json:"Facturado"`
	NumeroFactura      string `json:"Numero_Factura"`
	Asegurar           string `json:"Asegurar"`
	AvisoNum           string `json:"Aviso_Num"`
	Asegurado          string `json:"Asegurado"`
	AreaVarma          string `json:"Area VARMA"`
}

type Proveedores []struct {
	Proveedor                string `json:"Proveedor"`
	NifProveedor             string `json:"NIF_Proveedor"`
	Actividadproveedor       string `json:"ActividadProveedor"`
	Domicilioproveedor       string `json:"DomicilioProveedor"`
	PoblaciNproveedor        string `json:"PoblaciónProveedor"`
	Provinciaproveedor       string `json:"ProvinciaProveedor"`
	Codpostproveedor         string `json:"CodPostProveedor"`
	Telefonoproveedor        string `json:"TelefonoProveedor"`
	Personacontactoproveedor string `json:"PersonaContactoProveedor"`
	Faxproveedor             string `json:"FaxProveedor"`
	Telefonomovilproveedor2  string `json:"TelefonoMovilProveedor2"`
	Telefonomovilproveedor   string `json:"TelefonoMovilProveedor"`
	Zonatrabajo              string `json:"ZonaTrabajo"`
	Observaciones            string `json:"Observaciones"`
}

type Provincias []struct {
	Provincia string `json:"Provincia"`
}

type RelacionFacturas []struct {
	NumeroDeFactura string `json:"Numero_de_Factura"`
	FechaDeEmisiN   string `json:"Fecha_de_Emisión"`
	Cliente         string `json:"Cliente"`
	SeguroFactura   string `json:"Seguro_Factura"`
	TipoIva         string `json:"Tipo_IVA"`
	FormaPagoFra    string `json:"Forma_Pago_Fra"`
	Vencimiento     string `json:"Vencimiento"`
	Cobrado         string `json:"Cobrado"`
	FechaCobro      string `json:"Fecha_cobro"`
	Contabilizado   string `json:"Contabilizado"`
}

type Tarifas []struct {
	Tarifa        string `json:"Tarifa"`
	Observaciones string `json:"Observaciones"`
}
