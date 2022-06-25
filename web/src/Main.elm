port module Main exposing (..)

import Browser
import Dict exposing (Dict)
import Element as E exposing (Element)
import Element.Background as B
import Element.Input as I
import Html exposing (Html)
import Http
import Json.Decode as JD
import Json.Encode as JE
import Maybe exposing (withDefault)
import Svg exposing (Svg)
import Svg.Attributes as SvgA
import Svg.Events as SvgE
import Task


port receiveSocketMsg : (JD.Value -> msg) -> Sub msg


port robotsNewLocationsReceived : (String -> msg) -> Sub msg


port tasksReceived : (String -> msg) -> Sub msg


main =
    Browser.element
        { init = init
        , update = update
        , view = view
        , subscriptions = subscriptions
        }


type Movement
    = S
    | W
    | N
    | E


type RouteState
    = NotStarted
    | InProgress
    | Cancelled
    | Finished


type alias Robot =
    { name : String, x : Int, y : Int }


type alias MovementRequest =
    { moveSequences : List String
    }


encodeMovementRequest : MovementRequest -> JE.Value
encodeMovementRequest req =
    JE.object
        [ ( "moveSequences", JE.list JE.string req.moveSequences ) ]


robotDecoder : JD.Decoder Robot
robotDecoder =
    JD.map3
        Robot
        (JD.field "id" (JD.int |> JD.map String.fromInt))
        (JD.field "xPosition" JD.int)
        (JD.field "yPosition" JD.int)


robotsDecoder : JD.Decoder (List Robot)
robotsDecoder =
    JD.list robotDecoder


type alias Route =
    { id : String
    , movements : List Movement
    , state : RouteState
    }


type alias Routes =
    Dict String (List Route)


type alias Model =
    { robots : List Robot
    , selected : Maybe String
    , movements : Maybe String
    }


type alias Flags =
    ()


init : Flags -> ( Model, Cmd Msg )
init () =
    ( { robots = []
      , selected = Nothing
      , movements = Nothing
      }
    , Cmd.batch
        [ getRobotsList
        ]
    )


type Msg
    = NothingHappened
    | RobotsListArrived (Result Http.Error (List Robot))
    | RobotsUpdatedStr String
    | TasksUpdatedStr String
    | RobotSelected (Maybe String)
    | MovementInputChanged String
    | MovementInputDispatched String String


subscriptions : Model -> Sub Msg
subscriptions _ =
    Sub.batch
        [ robotsNewLocationsReceived RobotsUpdatedStr
        , tasksReceived TasksUpdatedStr
        ]


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        NothingHappened ->
            ( model, Cmd.none )

        RobotsListArrived robotRes ->
            case robotRes of
                Err e ->
                    ( model, Cmd.none )

                Ok r ->
                    ( { model | robots = model.robots ++ r }
                    , Cmd.none
                    )

        RobotsUpdatedStr v ->
            case JD.decodeString robotsDecoder v of
                Err e ->
                    ( model, Cmd.none )

                Ok r ->
                    ( { model | robots = updateRobotsMap model.robots r }
                    , Cmd.none
                    )

        TasksUpdatedStr v ->
            ( model, Cmd.none )

        RobotSelected name_ ->
            ( { model | selected = name_ }
            , Cmd.none
            )

        MovementInputChanged str ->
            let
                valid =
                    List.all (\c -> c == 'N' || c == 'S' || c == 'W' || c == 'E') (String.toList <| String.toUpper str)
            in
            if valid then
                ( { model | movements = Just <| String.toUpper str }, Cmd.none )

            else
                ( model, Cmd.none )

        MovementInputDispatched id seq ->
            ( { model | movements = Nothing }
            , moveRobot id seq
            )


updateRobotsMap : List Robot -> List Robot -> List Robot
updateRobotsMap l1 l2 =
    List.map (\i -> ( i.name, i )) l1
        ++ List.map (\i -> ( i.name, i )) l2
        |> Dict.fromList
        |> Dict.toList
        |> List.map (\p -> Tuple.second p)


view : Model -> Html Msg
view model =
    E.layout
        [ E.width E.fill, E.height E.fill ]
        (E.column
            [ E.centerX, E.centerY ]
            [ board model
            , E.row
                [ E.padding 5, E.spacing 5, E.centerX ]
                [ I.text []
                    { onChange = MovementInputChanged
                    , text = model.movements |> Maybe.withDefault ""
                    , placeholder = Nothing
                    , label = I.labelLeft [] (E.text "")
                    }
                , I.button [ E.padding 10, B.color <| E.rgb 0.5 0.8 0 ]
                    { onPress =
                        Just <|
                            MovementInputDispatched
                                (model.selected |> Maybe.withDefault "")
                                (model.movements |> Maybe.withDefault "")
                    , label = E.el [] (E.text "go")
                    }
                ]
            ]
        )


board model =
    let
        robots =
            model.robots |> List.map (\r -> robotFace model r)
    in
    E.html <|
        Svg.svg
            [ SvgA.width "400"
            , SvgA.height "400"
            , SvgA.viewBox "0 0 100 100"
            ]
            ([ gridBackground
             , gridVLines
             , gridHLines

             -- , robotFace <| Robot "A" 3 1
             ]
                ++ robots
            )


gridBackground : Svg Msg
gridBackground =
    Svg.rect
        [ SvgA.x "0"
        , SvgA.y "0"
        , SvgA.width "400"
        , SvgA.height "400"
        , SvgA.fill "#cdeef0"
        , SvgE.onClick <| RobotSelected Nothing
        ]
        []


gridVLines : Svg Msg
gridVLines =
    List.range 1 10
        |> List.map (\a -> a * 10)
        |> List.map
            (\a ->
                Svg.line
                    [ SvgA.x1 <| String.fromInt (a + 0)
                    , SvgA.x2 <| String.fromInt (a + 0)
                    , SvgA.y2 "0"
                    , SvgA.y2 "400"
                    , SvgA.stroke "black"
                    , SvgA.strokeWidth "0.2"
                    ]
                    []
            )
        |> Svg.svg []


gridHLines : Svg Msg
gridHLines =
    List.range 1 10
        |> List.map (\a -> a * 10)
        |> List.map
            (\a ->
                Svg.line
                    [ SvgA.y1 <| String.fromInt (a + 0)
                    , SvgA.y2 <| String.fromInt (a + 0)
                    , SvgA.x2 "0"
                    , SvgA.x2 "400"
                    , SvgA.stroke "black"
                    , SvgA.strokeWidth "0.2"
                    ]
                    []
            )
        |> Svg.svg []


robotFace : Model -> Robot -> Svg Msg
robotFace model r =
    let
        { name, x, y } =
            transformCoordinate r

        selected =
            model.selected |> withDefault ""
    in
    Svg.g
        [ SvgE.onClick <| RobotSelected <| Just name
        , SvgA.cursor "pointer"
        ]
        [ Svg.circle
            [ SvgA.cx <| String.fromInt <| (x * 10) + 5
            , SvgA.cy <| String.fromInt <| (y * 10) + 5
            , SvgA.r "4"
            , if selected == r.name then
                SvgA.fill "#ebb134"

              else
                SvgA.fill "#4f2bdf"
            ]
            []
        , Svg.text_
            [ SvgA.x <| String.fromFloat <| toFloat (x * 10) + 3.7
            , SvgA.y <| String.fromFloat <| toFloat (y * 10) + 6.3
            , SvgA.fill "#efedf5"
            , SvgA.fontSize "4"
            ]
            [ Svg.text name ]
        ]


transformCoordinate : Robot -> Robot
transformCoordinate r =
    let
        newY =
            modBy 10 r.y - 9 |> abs

        newX =
            r.x
    in
    { r | y = newY, x = newX }


getRobotsList : Cmd Msg
getRobotsList =
    Http.get
        { url = "/api/robots"
        , expect = Http.expectJson RobotsListArrived robotsDecoder
        }


moveRobot : String -> String -> Cmd Msg
moveRobot id seq =
    Http.request
        { method = "PUT"
        , headers = []
        , url = "/api/robots/" ++ id
        , body = Http.jsonBody <| encodeMovementRequest <| MovementRequest <| (seq |> String.toList |> List.map String.fromChar)
        , expect = Http.expectString (\a -> NothingHappened)
        , timeout = Nothing
        , tracker = Nothing
        }
