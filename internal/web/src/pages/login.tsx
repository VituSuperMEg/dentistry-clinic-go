import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { api } from "@/services/api";
import { Form, Formik } from "formik";

export const Login = () => {
  return (
    <Formik
      initialValues={{
        email: "",
        senha: "",
      }}
      onSubmit={async (values) => {
        api
          .post("/login", values)
          .then((response) => {
            console.log("Login bem-sucedido:", response.data);
          })
          .catch((error) => {
            console.error(
              "Erro no login:",
              error.response ? error.response.data : error.message
            );
          });
      }}
    >
      {({ handleChange }) => (
        <Form className="centered bg-white shadow-md h-[400px] w-[400px] rounded">
          <div className="flex flex-col p-10 gap-2 justify-between h-full">
            <header>
              <h1 className="text-3xl">
                Login<span className="text-indigo-500">.</span>
              </h1>
              <p className="text-sm text-zinc-500">
                Gerencie seus pacientes e seus dentes.
              </p>
            </header>
            <div className="flex gap-5 flex-col">
              <div>
                <Label>E-mail</Label>
                <Input type="text" name="email" onChange={handleChange} />
              </div>
              <div>
                <Label>Senha</Label>
                <Input type="password" name="senha" onChange={handleChange} />
              </div>
            </div>
            <Button type="submit" variant={"primary"}>
              Entrar
            </Button>
          </div>
        </Form>
      )}
    </Formik>
  );
};
