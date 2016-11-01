// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"DVEapi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/informacion_sociedad_participante",
			beego.NSInclude(
				&controllers.InformacionSociedadParticipanteController{},
			),
		),

		beego.NSNamespace("/informacion_proveedor",
			beego.NSInclude(
				&controllers.InformacionProveedorController{},
			),
		),

		beego.NSNamespace("/informacion_sociedad_temporal",
			beego.NSInclude(
				&controllers.InformacionSociedadTemporalController{},
			),
		),

		beego.NSNamespace("/ciiu_clase",
			beego.NSInclude(
				&controllers.CiiuClaseController{},
			),
		),

		beego.NSNamespace("/dependencia",
			beego.NSInclude(
				&controllers.DependenciaController{},
			),
		),

		beego.NSNamespace("/banco",
			beego.NSInclude(
				&controllers.BancoController{},
			),
		),

		beego.NSNamespace("/contrato",
			beego.NSInclude(
				&controllers.ContratoController{},
			),
		),

		beego.NSNamespace("/departamento",
			beego.NSInclude(
				&controllers.DepartamentoController{},
			),
		),

		beego.NSNamespace("/codigo_validacion",
			beego.NSInclude(
				&controllers.CodigoValidacionController{},
			),
		),

		beego.NSNamespace("/inhabilidad",
			beego.NSInclude(
				&controllers.InhabilidadController{},
			),
		),

		beego.NSNamespace("/informacion_persona_natural",
			beego.NSInclude(
				&controllers.InformacionPersonaNaturalController{},
			),
		),

		beego.NSNamespace("/ciudad",
			beego.NSInclude(
				&controllers.CiudadController{},
			),
		),

		beego.NSNamespace("/informacion_persona_juridica",
			beego.NSInclude(
				&controllers.InformacionPersonaJuridicaController{},
			),
		),

		beego.NSNamespace("/contrato_proveedor",
			beego.NSInclude(
				&controllers.ContratoProveedorController{},
			),
		),

		beego.NSNamespace("/evaluacion",
			beego.NSInclude(
				&controllers.EvaluacionController{},
			),
		),

		beego.NSNamespace("/ciiu_tipo",
			beego.NSInclude(
				&controllers.CiiuTipoController{},
			),
		),

		beego.NSNamespace("/menu",
			beego.NSInclude(
				&controllers.MenuController{},
			),
		),

		beego.NSNamespace("/nomenclatura_dian",
			beego.NSInclude(
				&controllers.NomenclaturaDianController{},
			),
		),

		beego.NSNamespace("/ciiu_division",
			beego.NSInclude(
				&controllers.CiiuDivisionController{},
			),
		),

		beego.NSNamespace("/ciiu_subclase",
			beego.NSInclude(
				&controllers.CiiuSubclaseController{},
			),
		),

		beego.NSNamespace("/ordenador_gasto",
			beego.NSInclude(
				&controllers.OrdenadorGastoController{},
			),
		),

		beego.NSNamespace("/rup_tipo",
			beego.NSInclude(
				&controllers.RupTipoController{},
			),
		),

		beego.NSNamespace("/objeto_contratar",
			beego.NSInclude(
				&controllers.ObjetoContratarController{},
			),
		),

		beego.NSNamespace("/supervisor",
			beego.NSInclude(
				&controllers.SupervisorController{},
			),
		),

		beego.NSNamespace("/pais",
			beego.NSInclude(
				&controllers.PaisController{},
			),
		),

		beego.NSNamespace("/snies_nucleo_basico",
			beego.NSInclude(
				&controllers.SniesNucleoBasicoController{},
			),
		),

		beego.NSNamespace("/proveedor_representante_legal",
			beego.NSInclude(
				&controllers.ProveedorRepresentanteLegalController{},
			),
		),

		beego.NSNamespace("/proveedor_actividad_ciiu",
			beego.NSInclude(
				&controllers.ProveedorActividadCiiuController{},
			),
		),

		beego.NSNamespace("/rup_especialidad",
			beego.NSInclude(
				&controllers.RupEspecialidadController{},
			),
		),

		beego.NSNamespace("/telefono",
			beego.NSInclude(
				&controllers.TelefonoController{},
			),
		),

		beego.NSNamespace("/parametro_estandar",
			beego.NSInclude(
				&controllers.ParametroEstandarController{},
			),
		),

		beego.NSNamespace("/proveedor_telefono",
			beego.NSInclude(
				&controllers.ProveedorTelefonoController{},
			),
		),

		beego.NSNamespace("/snies_area",
			beego.NSInclude(
				&controllers.SniesAreaController{},
			),
		),

		beego.NSNamespace("/solicitud_cotizacion",
			beego.NSInclude(
				&controllers.SolicitudCotizacionController{},
			),
		),

		beego.NSNamespace("/objeto_contratar_nucleo_basico",
			beego.NSInclude(
				&controllers.ObjetoContratarNucleoBasicoController{},
			),
		),

		beego.NSNamespace("/objeto_contratar_actividad_ciiu",
			beego.NSInclude(
				&controllers.ObjetoContratarActividadCiiuController{},
			),
		),

		beego.NSNamespace("/rup_grupo",
			beego.NSInclude(
				&controllers.RupGrupoController{},
			),
		),

		beego.NSNamespace("/unidad",
			beego.NSInclude(
				&controllers.UnidadController{},
			),
		),

		beego.NSNamespace("/tipo_conformacion",
			beego.NSInclude(
				&controllers.TipoConformacionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
