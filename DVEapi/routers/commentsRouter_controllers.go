package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["DVEapi/controllers:BancoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:BancoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:BancoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:BancoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:BancoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:BancoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:BancoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:BancoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:BancoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:BancoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CiiuClaseController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CiiuClaseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CiiuClaseController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CiiuClaseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CiiuClaseController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CiiuClaseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CiiuClaseController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CiiuClaseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CiiuClaseController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CiiuClaseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CiiuDivisionController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CiiuDivisionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CiiuDivisionController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CiiuDivisionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CiiuDivisionController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CiiuDivisionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CiiuDivisionController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CiiuDivisionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CiiuDivisionController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CiiuDivisionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CiiuSubclaseController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CiiuSubclaseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CiiuSubclaseController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CiiuSubclaseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CiiuSubclaseController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CiiuSubclaseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CiiuSubclaseController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CiiuSubclaseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CiiuSubclaseController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CiiuSubclaseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CiiuTipoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CiiuTipoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CiiuTipoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CiiuTipoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CiiuTipoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CiiuTipoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CiiuTipoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CiiuTipoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CiiuTipoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CiiuTipoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CiudadController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CiudadController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CiudadController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CiudadController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CiudadController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CiudadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CodigoValidacionController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CodigoValidacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CodigoValidacionController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CodigoValidacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CodigoValidacionController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CodigoValidacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CodigoValidacionController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CodigoValidacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:CodigoValidacionController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:CodigoValidacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ContratoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ContratoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ContratoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ContratoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ContratoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ContratoProveedorController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ContratoProveedorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ContratoProveedorController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ContratoProveedorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ContratoProveedorController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ContratoProveedorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ContratoProveedorController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ContratoProveedorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ContratoProveedorController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ContratoProveedorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InformacionPersonaJuridicaController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InformacionPersonaJuridicaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InformacionPersonaJuridicaController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InformacionPersonaJuridicaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InformacionPersonaJuridicaController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InformacionPersonaJuridicaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InformacionPersonaJuridicaController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InformacionPersonaJuridicaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InformacionPersonaJuridicaController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InformacionPersonaJuridicaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InformacionSociedadParticipanteController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InformacionSociedadParticipanteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InformacionSociedadParticipanteController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InformacionSociedadParticipanteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InformacionSociedadParticipanteController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InformacionSociedadParticipanteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InformacionSociedadParticipanteController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InformacionSociedadParticipanteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InformacionSociedadParticipanteController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InformacionSociedadParticipanteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InformacionSociedadTemporalController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InformacionSociedadTemporalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InformacionSociedadTemporalController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InformacionSociedadTemporalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InformacionSociedadTemporalController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InformacionSociedadTemporalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InformacionSociedadTemporalController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InformacionSociedadTemporalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InformacionSociedadTemporalController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InformacionSociedadTemporalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InhabilidadController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InhabilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InhabilidadController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InhabilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InhabilidadController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InhabilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InhabilidadController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InhabilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:InhabilidadController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:InhabilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:MenuController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:MenuController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:MenuController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:MenuController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:MenuController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:MenuController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:MenuController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:MenuController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:MenuController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:MenuController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:NomenclaturaDianController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:NomenclaturaDianController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:NomenclaturaDianController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:NomenclaturaDianController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:NomenclaturaDianController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:NomenclaturaDianController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:NomenclaturaDianController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:NomenclaturaDianController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:NomenclaturaDianController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:NomenclaturaDianController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarActividadCiiuController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarActividadCiiuController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarActividadCiiuController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarActividadCiiuController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarActividadCiiuController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarActividadCiiuController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarActividadCiiuController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarActividadCiiuController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarActividadCiiuController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarActividadCiiuController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarNucleoBasicoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarNucleoBasicoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarNucleoBasicoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarNucleoBasicoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarNucleoBasicoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarNucleoBasicoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarNucleoBasicoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ObjetoContratarNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:PaisController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:PaisController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:PaisController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:PaisController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:PaisController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:PaisController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:PaisController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:PaisController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:PaisController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:PaisController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ProveedorActividadCiiuController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ProveedorActividadCiiuController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ProveedorActividadCiiuController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ProveedorActividadCiiuController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ProveedorActividadCiiuController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ProveedorActividadCiiuController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ProveedorActividadCiiuController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ProveedorActividadCiiuController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ProveedorActividadCiiuController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ProveedorActividadCiiuController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ProveedorRepresentanteLegalController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ProveedorRepresentanteLegalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ProveedorRepresentanteLegalController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ProveedorRepresentanteLegalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ProveedorRepresentanteLegalController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ProveedorRepresentanteLegalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ProveedorRepresentanteLegalController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ProveedorRepresentanteLegalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ProveedorRepresentanteLegalController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ProveedorRepresentanteLegalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ProveedorTelefonoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ProveedorTelefonoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ProveedorTelefonoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ProveedorTelefonoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ProveedorTelefonoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ProveedorTelefonoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ProveedorTelefonoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ProveedorTelefonoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:ProveedorTelefonoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:ProveedorTelefonoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:RupEspecialidadController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:RupEspecialidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:RupEspecialidadController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:RupEspecialidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:RupEspecialidadController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:RupEspecialidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:RupEspecialidadController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:RupEspecialidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:RupEspecialidadController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:RupEspecialidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:RupGrupoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:RupGrupoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:RupGrupoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:RupGrupoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:RupGrupoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:RupGrupoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:RupGrupoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:RupGrupoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:RupGrupoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:RupGrupoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:RupTipoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:RupTipoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:RupTipoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:RupTipoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:RupTipoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:RupTipoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:RupTipoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:RupTipoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:RupTipoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:RupTipoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:SniesAreaController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:SniesAreaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:SniesAreaController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:SniesAreaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:SniesAreaController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:SniesAreaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:SniesAreaController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:SniesAreaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:SniesAreaController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:SniesAreaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:SolicitudCotizacionController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:SolicitudCotizacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:SolicitudCotizacionController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:SolicitudCotizacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:SolicitudCotizacionController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:SolicitudCotizacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:SolicitudCotizacionController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:SolicitudCotizacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:SolicitudCotizacionController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:SolicitudCotizacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:SupervisorController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:SupervisorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:SupervisorController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:SupervisorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:SupervisorController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:SupervisorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:SupervisorController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:SupervisorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:SupervisorController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:SupervisorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:TelefonoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:TelefonoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:TelefonoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:TelefonoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:TelefonoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:TelefonoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:TelefonoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:TelefonoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:TelefonoController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:TelefonoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:TipoConformacionController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:TipoConformacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:TipoConformacionController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:TipoConformacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:TipoConformacionController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:TipoConformacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:TipoConformacionController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:TipoConformacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:TipoConformacionController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:TipoConformacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:UnidadController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:UnidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:UnidadController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:UnidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:UnidadController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:UnidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:UnidadController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:UnidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["DVEapi/controllers:UnidadController"] = append(beego.GlobalControllerRouter["DVEapi/controllers:UnidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
