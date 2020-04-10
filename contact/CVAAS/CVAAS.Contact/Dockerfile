#See https://aka.ms/containerfastmode to understand how Visual Studio uses this Dockerfile to build your images for faster debugging.

FROM mcr.microsoft.com/dotnet/core/aspnet:3.1-buster-slim AS base
WORKDIR /app
EXPOSE 80
EXPOSE 443

FROM mcr.microsoft.com/dotnet/core/sdk:3.1-buster AS build
WORKDIR /src
COPY ["CVAAS.Contact/CVAAS.Contact.csproj", "CVAAS.Contact/"]
RUN dotnet restore "CVAAS.Contact/CVAAS.Contact.csproj"
COPY . .
WORKDIR "/src/CVAAS.Contact"
RUN dotnet build "CVAAS.Contact.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "CVAAS.Contact.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "CVAAS.Contact.dll"]