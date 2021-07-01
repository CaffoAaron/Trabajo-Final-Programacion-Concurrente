import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import React, { useState } from "react";
import { faPen } from "@fortawesome/free-solid-svg-icons";
import {
  Card,
  CardBody,
  CardTitle,
  FormGroup,
  Label,
  Input,
  Row,
  Col,

} from "reactstrap";
import * as Yup from "yup";
import { Formik, Form } from "formik";
import Modal from "../components/modal";
import axios from "axios";

export default function FormView() {
  const options = ["Selecciona una respuesta","No", "Si"];
  //formSchema se encarga de las validaciones y tipo de dato de los campos
  const formSchema = Yup.object().shape({
    casado: Yup.string().required("Seleccione una respuesta"),
    hijos: Yup.string().required("Seleccione una respuesta"),
    carrera_universitaria: Yup.string().required("Seleccione una respuesta"),
    casa_propia: Yup.string().required("Seleccione una respuesta"),
    otro_prestamo: Yup.string().required("Seleccione una respuesta"),
  });
  const [disabledSubmit, setDisabledSubmit] = useState(true);

  /*async function onSubmit(values) {
    const payload = {
      ...values,
      condiciones: !disabledSubmit,
    };
    console.log(payload);
  }*/

  return (
    <Card>
      <CardBody>
        <span className="center padd-bot-3">
          {/*Colocamos un ícono importado desde la librería de FontAwesome*/}
          <FontAwesomeIcon icon={faPen} className="icon" />
        </span>
        <CardTitle tag="h5" className="center">
          Solicitud prestamo Peru independiente
        </CardTitle>

        <Formik
          /*Se Inicializan los valores de los campos del formulario y se declara el objeto a
          validar en el formulario*/
          enableReinitialize={true}
          initialValues={{
            casado: "",
              mensaje: "",
            hijos: "",
            carrera_universitaria: "",
            casa_propia: "",
            otro_prestamo: "",
            mas_de_4_Años_como_empresa: "",
            mas_de_1_Local: "",
            mas_de_10_Empleados: "",
            Pago_de_Igv_Ultimos_6_Meses: "",
            declaron_confidencial_patrimonio: "",
          }}
          validationSchema={formSchema}
          //onSubmit={onSubmit}
        >
          {({ errors, touched, handleChange, handleBlur, values }) => (
            <Form className="padd-top-10">
              <Row form>
                <Col md={12}>
                  <FormGroup>
                    <Label>¿Su estado civil es casado ?</Label>
                    <select
                        name="casado"
                        className="form-control"
                        value={values.casado}
                        onChange={handleChange}
                        onBlur={handleBlur}
                    >
                      {options.map((option, index) => (
                          <option key={index} value={option}>
                            {option}
                          </option>
                      ))}
                    </select>
                  </FormGroup>
                </Col>
                <Col md={12}>
                  <FormGroup>
                    <Label>¿Tiene hijos?</Label>
                    <select
                        name="hijos"
                        className="form-control"
                        value={values.hijos}
                        onChange={handleChange}
                        onBlur={handleBlur}
                    >
                      {options.map((option, index) => (
                          <option key={index} value={option}>
                            {option}
                          </option>
                      ))}
                    </select>
                  </FormGroup>
                </Col>
              </Row>
              <Row form>
                <Col md={12}>
                  <FormGroup>
                    <Label>¿Cuenta con carrera universitaria?</Label>
                    <select
                        name="carrera_universitaria"
                        className="form-control"
                        value={values.carrera_universitaria}
                        onChange={handleChange}
                        onBlur={handleBlur}
                    >
                      {options.map((option, index) => (
                          <option key={index} value={option}>
                            {option}
                          </option>
                      ))}
                    </select>
                  </FormGroup>
                </Col>
              </Row>
              <Row form>
                <Col md={12}>
                  <FormGroup>
                    <Label>¿Cuenta con casa propia?</Label>
                    <select
                      name="casa_propia"
                      className="form-control"
                      value={values.casa_propia}
                      onChange={handleChange}
                      onBlur={handleBlur}
                    >
                      {options.map((option, index) => (
                        <option key={index} value={option}>
                          {option}
                        </option>
                      ))}
                    </select>
                  </FormGroup>
                </Col>
                <Col md={12}>
                  <FormGroup>
                    <Label>¿Cuenta con alguin prestamo en algun banco?</Label>
                    <select
                        name="otro_prestamo"
                        className="form-control"
                        value={values.otro_prestamo}
                        onChange={handleChange}
                        onBlur={handleBlur}
                    >
                      {options.map((option, index) => (
                          <option key={index} value={option}>
                            {option}
                          </option>
                      ))}
                    </select>
                  </FormGroup>
                </Col>
                <Col md={12}>
                  <FormGroup>
                    <Label>¿Su empresa cuanta con mas de 4 año de creada?</Label>
                    <select
                        name="mas_de_4_Años_como_empresa"
                        className="form-control"
                        value={values.mas_de_4_Años_como_empresa}
                        onChange={handleChange}
                        onBlur={handleBlur}
                    >
                      {options.map((option, index) => (
                          <option key={index} value={option}>
                            {option}
                          </option>
                      ))}
                    </select>
                  </FormGroup>
                </Col>
                <Col md={12}>
                  <FormGroup>
                    <Label>¿Cuenta con mas de un de local?</Label>
                    <select
                        name="mas_de_1_Local"
                        className="form-control"
                        value={values.mas_de_1_Local}
                        onChange={handleChange}
                        onBlur={handleBlur}
                    >
                      {options.map((option, index) => (
                          <option key={index} value={option}>
                            {option}
                          </option>
                      ))}
                    </select>
                  </FormGroup>
                </Col>
                <Col md={12}>
                  <FormGroup>
                    <Label>¿Cuenta con mas de 10 empreados contrado al momento de hoy ?</Label>
                    <select
                        name="mas_de_10_Empleados"
                        className="form-control"
                        value={values.mas_de_10_Empleados}
                        onChange={handleChange}
                        onBlur={handleBlur}
                    >
                      {options.map((option, index) => (
                          <option key={index} value={option}>
                            {option}
                          </option>
                      ))}
                    </select>
                  </FormGroup>
                </Col>
                <Col md={12}>
                  <FormGroup>
                    <Label>¿A realizado el pago de su igv en los ultimos 6 meses?</Label>
                    <select
                        name="Pago_de_Igv_Ultimos_6_Meses"
                        className="form-control"
                        value={values.Pago_de_Igv_Ultimos_6_Meses}
                        onChange={handleChange}
                        onBlur={handleBlur}
                    >
                      {options.map((option, index) => (
                          <option key={index} value={option}>
                            {option}
                          </option>
                      ))}
                    </select>
                  </FormGroup>
                </Col>
                <Col md={12}>
                  <FormGroup>
                    <Label>¿A realizado la declaracion confidencial del patrimonio?</Label>
                    <select
                        name="declaron_confidencial_patrimonio"
                        className="form-control"
                        value={values.declaron_confidencial_patrimonio}
                        onChange={handleChange}
                        onBlur={handleBlur}
                    >
                      {options.map((option, index) => (
                          <option key={index} value={option}>
                            {option}
                          </option>
                      ))}
                    </select>
                  </FormGroup>
                </Col>
              </Row>
              <FormGroup check className="padd-bot-5">
                {/*En el evento onClick ejecutaremos una función que cambiara el estado de los botones y del checkbox*/}
                <Input
                  type="checkbox"
                  onClick={() => {
                      setDisabledSubmit(!disabledSubmit);
                      axios.post(`http://localhost:9200/knn?casado=`+ values.casado + `&hijos=`
                          +values.hijos +`&carrera_universitaria=` + values.carrera_universitaria + `&casa_propia=`+ values.casa_propia
                          + `&otro_prestamo=`+ values.otro_prestamo +`&mas_de_4_Años_como_empresa=` + values.mas_de_4_Años_como_empresa
                          + `&mas_de_1_Local=`+ values.mas_de_1_Local + `&mas_de_10_Empleados=` + values.mas_de_10_Empleados
                          + `&Pago_de_Igv_Ultimos_6_Meses=`+ values.Pago_de_Igv_Ultimos_6_Meses
                          +`&declaron_confidencial_patrimonio=` + values.declaron_confidencial_patrimonio
                          , {
                              "casado": values.casado,
                              "Hijos": values.hijos,} )
                          .then(res => {
                              values.mensaje=res.data.Mensaje;
                          })
                  }}
                />
                <Label check >Acepto los terminos y condiciones</Label>
              </FormGroup>

              {/*<Button
                block
                outline
                color="primary"
                disabled={disabledSubmit}
                type="submit"
              >
                Enviar
              </Button>*/}
              {/*Enviamos al modal los props correspondiente al estado, valores y errores*/}
              <Modal
                values={values}
                errors={errors}
                disabledSubmit={disabledSubmit}
              />
            </Form>
          )}
        </Formik>
      </CardBody>
    </Card>
  );
}
