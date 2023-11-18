export const urlParser = (protocol :"WS" | "HTTP" ,path:string,isProjectorToken?:boolean)=>{
    let urlPath = path
    if(urlPath.startsWith("\/")){
       urlPath = urlPath.substring(1)
    }
    if(urlPath.endsWith("\/")){
       urlPath = urlPath.substring(0,urlPath.length-1)
    }
    const tokenQuery = new URLSearchParams({
        token: isProjectorToken?  import.meta.env.VITE_PROJECTOR_TOKEN :""
    })
    return `${protocol === "WS" ?"ws://" : "http://"}${import.meta.env.VITE_BASE_URL}/${urlPath}${isProjectorToken ? "?"+ tokenQuery.toString() : ""}`
}