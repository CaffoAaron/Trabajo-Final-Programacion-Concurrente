import React, { useState } from "react";
import { Button, Modal, ModalHeader, ModalBody, ModalFooter } from "reactstrap";


export default function ModalComponent({ values, errors, disabledSubmit }) {
  const [modal, setModal] = useState(false);
  const toggle = () => setModal(!modal);
  const reload = () => (window.location = "/");


    return (
    <>
      <Button
        block
        outline
        color="primary"
        onClick={toggle}
        type="submit"
        disabled={disabledSubmit}
      >
        Enviar
      </Button>

      <Modal isOpen={modal} toggle={toggle}>
        <ModalHeader toggle={toggle}>Resultado</ModalHeader>
        {/*En el cuerpo del modal vamos a mostrar la información en formato JSON de los valores y errores acumulados*/}
        <ModalBody>
          <h5 >{JSON.stringify(values.mensaje, null, 2)}</h5>

        </ModalBody>
        <ModalFooter>
          {/*Al presionar el botón recargar vamos a actualizar la página y borrar el formulario*/}
          <Button color="primary" onClick={reload}>
            Volver
          </Button>
        </ModalFooter>
      </Modal>
    </>
  );
}
