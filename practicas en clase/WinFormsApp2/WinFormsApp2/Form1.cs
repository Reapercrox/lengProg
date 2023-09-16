using ClassLibrary1;
using Microsoft.FSharp.Collections;
using Microsoft.FSharp.Core;

namespace WinFormsApp2
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void button1_Click(object sender, EventArgs e)
        {
            if (textBox1.Text != "")
            {
                var caracteres = textBox1.Text.ToCharArray();
                //var listaDeCaracteres = new FSharpList<char>(caracteres[0], caracteres);
                var secuenciaFSharp = ListModule.OfSeq(caracteres);
                var listaDeCaracteres = ListModule.OfSeq(secuenciaFSharp);

                var listaresultado = ClassLibrary1.lib.invertirString(listaDeCaracteres);

                textBox2.Text = listaresultado. ToString(); 
            }
            else
                textBox2.Text = "Cannot convert empty String";
        }
    }
}