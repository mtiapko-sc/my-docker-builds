var builder = WebApplication.CreateBuilder(args);
builder.WebHost.UseUrls("http://0.0.0.0:8080");
var app = builder.Build();
app.MapGet("/healthz", () => "ok");
app.MapGet("/", () => "Hello from .NET!");
app.Run();
