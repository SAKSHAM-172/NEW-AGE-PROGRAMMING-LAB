open System

type Shape(name: string) =
    member this.Name = name
    abstract Area : unit -> float
    default this.Area() = 0.0
    abstract Describe : unit -> string
    default this.Describe() = sprintf "%s with area %.2f" name (this.Area())

type Circle(radius: float) =
    inherit Shape("Circle")
    override this.Area() = Math.PI * radius * radius

type Rectangle(width: float, height: float) =
    inherit Shape("Rectangle")
    override this.Area() = width * height

type Triangle(baseLen: float, height: float) =
    inherit Shape("Triangle")
    override this.Area() = 0.5 * baseLen * height

let baseShape = Shape("Generic Shape")
let circle = Circle(5.0)
let rectangle = Rectangle(4.0, 6.0)
let triangle = Triangle(3.0, 4.0)

printfn "=== Calling methods on each object directly ==="
printfn "%s area: %.2f" baseShape.Name (baseShape.Area())
printfn "%s area: %.2f" circle.Name (circle.Area())
printfn "%s area: %.2f" rectangle.Name (rectangle.Area())
printfn "%s area: %.2f" triangle.Name (triangle.Area())

printfn "\n=== Polymorphism: one list of Shape, different behaviour ==="
let shapes : Shape list =
    [ baseShape
      circle :> Shape
      rectangle :> Shape
      triangle :> Shape
      Circle(2.5) :> Shape
      Rectangle(3.0, 3.0) :> Shape ]

for s in shapes do
    printfn "%s" (s.Describe())

let totalArea = shapes |> List.sumBy (fun s -> s.Area())
printfn "\nTotal area of all shapes: %.2f" totalArea