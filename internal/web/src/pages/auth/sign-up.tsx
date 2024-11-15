import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { api } from "@/services/api";
import { Formik } from "formik";
import { useNavigate } from "react-router-dom";
import Swal from "sweetalert2";

export function SignUp() {
  const navigate = useNavigate();
  return (
    <Formik
      initialValues={{
        email: "",
        senha: "",
      }}
      onSubmit={async (values) => {
        await api
          .post("/login", values)
          .then(() => {
            navigate("/dashboard/home");
          })
          .catch(() => {
            Swal.fire({
              title: "",
              text: "E-mail ou senha inválido",
              icon: "error",
            });
          });
      }}
    >
      {({ submitForm, handleChange }) => (
        <section className="flex justify-center items-center h-screen">
          <div className="flex flex-col items-center justify-center bg-white shadow rounded p-10 w-[500px]">
            <div className="w-full">
              <h2 className="font-bold">Login</h2>
              <p className="text-sm font-normal text-zinc-500">
                Controle seus pacientes e suas evoluções.
              </p>
            </div>
            <form className="mt-8 mb-2 mx-auto flex flex-col gap-2 w-full">
              <div className="mb-1 flex flex-col gap-6">
                <div>
                  <Label>E-mail</Label>
                  <Input
                    name="email"
                    type="text"
                    placeholder="name@mail.com"
                    className=" !border-t-blue-gray-200 focus:!border-t-gray-900"
                    onChange={handleChange}
                  />
                </div>
                <div>
                  <Label>Senha</Label>
                  <Input
                    name="senha"
                    type="password"
                    placeholder="***"
                    className=" !border-t-blue-gray-200 focus:!border-t-gray-900"
                    onChange={handleChange}
                  />
                </div>
              </div>
              <Button
                variant="primary"
                type="button"
                className="mt-6"
                onClick={submitForm}
              >
                Entrar
              </Button>
            </form>
          </div>
        </section>
      )}
    </Formik>
  );
}

export default SignUp;
