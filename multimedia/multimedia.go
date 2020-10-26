package multimedia

type Imagen struct{

	Titulo string
	Formato string
	Canales string

}

func (i *Imagen) Mostrar() string{
	return "\nTitulo : " + i.Titulo +
		   "\nFormato: " + i.Formato +
		   "\nCanales: " + i.Canales +
		   "\n"
}

type Audio struct{
	
	Titulo     string
 	Formato    string
	Duracion   string

}

func (a *Audio) Mostrar() string{
	return "\nTitulo : " + a.Titulo +
		   "\nFormato: " + a.Formato +
		   "\nDuracion: " + a.Duracion +
		   "\n"
}

type Video struct{

	Titulo    string
	Formato   string
	Frames    string

}

func (v *Video) Mostrar() string{
	return "\nTitulo : " + v.Titulo +
		   "\nFormato: " + v.Formato +
		   "\nFrames: " + v.Frames +
		   "\n"
}

type Multimedia interface {
	Mostrar()string
}



type ContenidoWeb struct{
	Multi []Multimedia  
}

func (c *ContenidoWeb) Mostrar() string {

	var slide string

	for _, s := range c.Multi {
		slide += s.Mostrar()
	}
	return slide
}