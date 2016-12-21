var data = [];
$.getJSON("resolucion.json", function(datos) {
$.each(datos["contratos"], function(idx,contrato) {
data[idx]=[contrato["nombre"], contrato["cedula"], contrato["expedida"], contrato["categoria"], contrato["dedicacion"], contrato["horas"], contrato["semanas"], contrato["valor"], contrato["disponibilidad"]];
});
});
function generar() {

var texto = "<div class='col-lg-12'><div class='panel panel-default'><div class='panel-heading'><h4 class='intro-text text-center'> <strong>Resoluci&oacute;n</strong></h4></div><div class='panel-body' id='pdf'></div><div class='panel-body'><button type='submit' class='form-control' onclick='descargar()'>Descargar resoluci&oacute;n</button></div></div></div>";
            $("#carta").html(texto);
            texto = "<iframe id='cartapdf' frameborder='0' width='100%' height='400'></iframe>";
            $("#pdf").append(texto);


var columns = ["Nombre", "Cédula", "Expedida", "Categoría", "Dedicación", "Horas semanales", "Semanas contratadas", "Valor contrato", "Disponibilidad"];

 var doc = new jsPDF('p', 'pt');

 doc.text("Resolución No 0224", 220, 60);
 doc.setFontSize(12);
 doc.text("CONSIDERANDO", 240, 80);
 doc.setFontSize(10);
 var text = doc.splitTextToSize("Que la Universidad Distrital Francisco José de aldas Puede Vincular docentes en las modalidades de Hora Cátedra, Medio Tiempo Ocasional y Tiempo Competo Ocasional, en virtud del articulo 13 del Acuerdo 011 de noviembre 15 de 2002" + '.', doc.internal.pageSize.width - 80, {});
 doc.text(text, 40, 100);

 text = doc.splitTextToSize("Que mediante decreto 1279 del 19 de junio de 2002, el Gobierno Nacional estableció un nuevo régimen salarial y prestacional de los docentes de las universidades estatales u oficiales del Orden Nacional, Departamental, Municipal y Distrital. (Artículos 3 y 4)" + '.', doc.internal.pageSize.width - 80, {});
 doc.text(text, 40, 140);

 text = doc.splitTextToSize("Que mediante el acuerdo 012 de noviembre 15 de 2002, se fijan los factores para el reconocimiento y pago de la Hora Cátedra para los docentes que prestan servicios en la modalidad de vinculación especial en los programas de Pregrado" + '.', doc.internal.pageSize.width - 80, {});
 doc.text(text, 40, 180);

 text = doc.splitTextToSize("Que mlos servicios de los docentes de vinculación especial señalados en el acuerdo 011 de noviembre 15 de 2002 expedido por el Consejo Superior Universitario, deberán ser reconocidos mediante Resolución, (Resoluciones 0013 de enero 31 de 2003, 0013-A de enero 31 de 2003, Ley 30 de 1992 y Acuerdo 003 de 1997, articulo 49 y Ley 4 de 1992)" + '.', doc.internal.pageSize.width - 80, {});
 doc.text(text, 40, 220);

 text = doc.splitTextToSize("Que para efectos de pago y liquidación el mes comprenderá (4) semanas o (30) dias laborales" + '.', doc.internal.pageSize.width - 80, {});
 doc.text(text, 40, 270);

 text = doc.splitTextToSize("Que la presente vinculación será de 4,25 meses que equivale a 17 semanas, y se reconocerá en mese de (4) semanas, iniciando en la segunda semana de febrero de 2013, será de la siguiente forma: TRES SEMANAS DE FEBRERO, (cuatro semanas) MARZO, (cuatro semanas) ABRIL, (cuatro semanas) MAYO Y (dos semanas) JUNIO de 2017" + '.', doc.internal.pageSize.width - 80, {});
 doc.text(text, 40, 295);

 doc.setFontSize(12);
 doc.text("RESUELVE", 245, 340);
 doc.setFontSize(10);

 text = doc.splitTextToSize("ARTÍCULO PRIMERO - adicionar horas lectivas y el valor que corresponda para el PRIMER PERIODO ACÁDEMICO del 2013, a los docentes de vinculación Especial en el esalafón y dedicación establecidas en la siguiente tabla" + '.', doc.internal.pageSize.width - 80, {});
 doc.text(text, 40, 370);

doc.autoTable(columns, data, {
                startY: 410,
                margin: {horizontal: 10},
                columnStyles: {text: {columnWidth: 250}},
                pageBreak: 'avoid',
                headerStyles: {rowHeight: 20, fontSize: 7},
                bodyStyles: {rowHeight: 12, fontSize: 7}
            });

 text = doc.splitTextToSize("ARTÍCULO SEGUNDO - El pago de los servcios prestados por los profesores de vinculación especial según su escalafón, se cancelará previa certificación de las horas efectivamente dictadas, expedida por el decano" + '.', doc.internal.pageSize.width - 80, {});
 doc.text(text, 40, 580);

 text = doc.splitTextToSize("PARÁGRAFO - El valor del punto en pesos para el reconocimiento y pago de los docentes de hora cátedra sera el que fije el gobierno nacional, mediante decreto cada año y que la universidad acoja mediante acto administrativo para los docentes de vinculación especial" + '.', doc.internal.pageSize.width - 80, {});
 doc.text(text, 40, 620);

 text = doc.splitTextToSize("ARTÍCULO TERCERO - El docente deberá cumplir con lasobligaciones inherentes a la naturaleza del servicio, contempladas en la ley, en los reglamentos de la universidad y en losplanes de trabajo entregados por el profesor y aprobados por el decano" + '.', doc.internal.pageSize.width - 80, {});
 doc.text(text, 40, 660);

 text = doc.splitTextToSize("PARÁGRAFO - En caso de incumplimiento o retiro del docente, la universidad mediante acto administrativo hará la liquidación con corte ala fecha del cumplido expedido por el decano" + '.', doc.internal.pageSize.width - 80, {});
 doc.text(text, 40, 700);

 text = doc.splitTextToSize("ARTÍCULO CUARTO - El gasto que ocaciona la presente resolución se hará con cargo al presupuesto de la actual vigencia previa certificación de disponibilidad presupuestal" + '.', doc.internal.pageSize.width - 80, {});
 doc.text(text, 40, 740);

 text = doc.splitTextToSize("PARÁGRAFO - En todo caso, los pagos correspondetes estarán sujetos a las apropiaciones presupuestales y a las transferencias de la secretaría de hacienda distrital" + '.', doc.internal.pageSize.width - 80, {});
 doc.text(text, 40, 780);
 
 text = doc.splitTextToSize("ARTÍCULO QUINTO - Comunicar la presente  resolución al Docente, quien  manifestará bajo la gravedad de juramento que no se encuentra incurso en inhabilidades o incompatibilidades de ley para acpetar la presente vinculación especial" + '.', doc.internal.pageSize.width - 80, {});
 doc.text(text, 40, 820);

 text = doc.splitTextToSize("ARTÍCULO SEXTO - La presente resolución rige a partir de la fecha de expedición y surte efectos a partir de la -segunda semana- de -FEBRERO de 2017- para el -primer periodo académico de 2017-, a los -06 días- de mes de -febrero de 2017-" + '.', doc.internal.pageSize.width - 80, {});
 doc.text(text, 40, 860);

 doc.setProperties({
  title: 'Resolución'
 });
 
 var string = doc.output('datauristring');
 $('#cartapdf').attr('src', string);
}