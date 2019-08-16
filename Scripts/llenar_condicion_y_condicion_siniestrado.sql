INSERT INTO CONDICION (id,nombre,codigo,activo)
VALUES
    (-1,'Jubilado Decreto N° 894/01 y/o Dec 2288/02','0',1),
    (-2,'SERVICIOS COMUNES Mayor de 18 años','1',1),
    (-3,'Jubilado','2',1),
    (-4,'Menor','3',1),
    (-5,'Menor Anterior','4',1),
    (-6,'SERVICIOS DIFERENCIADOS Mayor de 18 años','5',1),
    (-7,'Pre- jubilables Sin relacion de dependencia -Sin servicios reales','6',1),
    (-8,'MEDIDA DE NO INNOVAR SERV. COMUNES','7',1),
    (-9,'MEDIDA DE NO INNOVAR SERV. DIFERENCIAD','8',1),
    (-10,'Jubilado Decreto N° 206/00 y/o Decreto Nº 894/01','9',1),
    (-11,'Pensión (NO SIPA)','10',1),
    (-12,'Pensión no Contributiva (NO SIPA)','11',1),
    (-13,'Art. 8º Ley Nº 27426','12',1),
    (-14,'Servicios Diferenciados no alcanzados por el Dto. 633/2018','13',1);

    DELETE FROM CONDICION WHERE ID = 1;

INSERT INTO CONDICIONSINIESTRADO(id,nombre,codigo,activo)
VALUES
(-1,'No Incapacitado','0',1),
(-2,'ILT Incapacidad Laboral Temporaria','1',1),
(-3,'LPPP Incapacidad Laboral Permanente Parcial Provisoria','2',1),
(-4,'ILPPD Incapacidad Laboral Permanente Parcial Definitiva','3',1),
(-5,'ILPTP Incapacidad Laboral Permanente Total Provisoria','4',1),
(-6,'Capital de recomposición Art. 15, ap. 3, Ley 24557','5',1),
(-7,'Ajuste Definitivo ILPPD de pago mensual','6',1),
(-8,'RENTA PERIODICA ILPPD Inc Lab Perm Parc Def  >50%<66%','7',1),
(-9,'SRT/SSN F.Garantía/F Reserva  ILT Incapacidad Laboral Temporaria','8',1), 
(-10,'SRT/SSN F.Garantía/F Reserva ILPPP Inc Lab Perm Parc Prov','9',1),
(-11,'SRT/SSN F.Garantía/F Reserva ILPTP Inc Lab Perm Total Prov','10',1),
(-12,'SRT/SSN F.Garantía/F Reserva ILPPD Inc Laboral Perm Parc Definitiva','11',1),
(-13,'ILPPD Beneficios devengados art. 11 p.4','12',1),
(-14,'INFORME Incremento salarial de trabajador siniestrado a ART','13',1);
